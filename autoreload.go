package main

import (
	"crypto/sha3"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type fileStamp struct {
	path    string
	hash    FileHash
	lastMod time.Time
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
var fsCache = make(map[string]fileStamp)

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
		fsCache[path] = fileStamp{
			path:    path,
			hash:    fhash,
			lastMod: ftime,
		}
		fsCacheMtx.Unlock()
	} else {
		if ftime.Equal(fstamp.lastMod) {
			return false
		}
		// note: we only calculate hash if timestamp check failed to avoid unnecessary computation
		fhash, err := getFileHash(path)
		if err != nil || fhash == fstamp.hash {
			return false
		}
		fsCacheMtx.Lock()
		fsCache[path] = fileStamp{
			path:    path,
			hash:    fhash,
			lastMod: ftime,
		}
		fsCacheMtx.Unlock()
	}

	return true
}

// Returns a list of file paths of files that were changed in dirPath with the specified extensions.
// The extensions must contain a starting dot like `.txt`
func DetectDirChanged(dirPath string, exts []string) []string {
	var changedFiles []string
	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
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
