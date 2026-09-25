package main

import (
	"crypto/sha3"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileStamp struct {
	Path    string    `json:"path"`
	Hash    FileHash  `json:"hash"`
	LastMod time.Time `json:"last_mod"`
}

type FileHash = string

func getFileHash(path string) (FileHash, error) {
	fContents, err := os.Open(path)
	if err != nil {
		return "", err
	}
	hasher := sha3.New256()
	if _, err := io.Copy(hasher, fContents); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

var fsCacheMtx sync.Mutex
var fsCache = make(map[string]FileStamp)

// detect is pull-based. every time you call if, it remembers the state of the
// file it was called with, so next time you call it, it will tell you if the contents of this file
// has changed or not since then.
func DetectFileChanged(path string) bool {
	fstat, err := os.Stat(path)
	if err != nil {
		return false
	}

	ftime := fstat.ModTime()
	fstamp, ok := fsCache[path]

	if !ok {
		fhash, err := getFileHash(path)
		if err != nil {
			return false
		}
		fsCacheMtx.Lock()
		fsCache[path] = FileStamp{
			Path:    path,
			Hash:    fhash,
			LastMod: ftime,
		}
		fsCacheMtx.Unlock()
	} else {
		if ftime.Equal(fstamp.LastMod) {
			return false
		}
		// note: we only calculate hash if timestamp check failed to avoid unnecessary computation
		fhash, err := getFileHash(path)
		if err != nil || fhash == fstamp.Hash {
			return false
		}
		fsCacheMtx.Lock()
		fsCache[path] = FileStamp{
			Path:    path,
			Hash:    fhash,
			LastMod: ftime,
		}
		fsCacheMtx.Unlock()
	}

	return true
}

// Returns a list of file paths of files that were changed in dirPath with the specified extensions.
// The extensions must contain a starting dot like `.txt`
// If no extensions are supplied every file will be checked
func DetectDirFilesChanged(dirPath string, exts ...string) []string {
	var changedFiles []string
	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		if len(exts) == 0 && DetectFileChanged(path) {
			changedFiles = append(changedFiles, path)
			return nil
		}

		ext := filepath.Ext(path)
		for _, extWanted := range exts {
			if ext == extWanted && DetectFileChanged(path) {
				changedFiles = append(changedFiles, path)
				break
			}
		}

		return nil
	})

	return changedFiles
}

func DumpFileHashes(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	fsCacheMtx.Lock()
	dump, err := json.Marshal(fsCache)
	fsCacheMtx.Unlock()
	if err != nil {
		return err
	}

	if _, err = f.Write(dump); err != nil {
		return err
	}
	return nil
}

func LoadFileHashes(path string) error {
	hashes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fsCacheMtx.Lock()
	err = json.Unmarshal(hashes, &fsCache)
	fsCacheMtx.Unlock()
	if err != nil {
		return err
	}

	return nil
}
