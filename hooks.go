package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

func CheckLinkIsReachable(linkpath string, node ast.Node) error {
	linkNode, ok := node.(*ast.Link)
	if !ok {
		return nil
	}

	link := string(linkNode.Destination)
	if !strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") {
		return nil
	}

	defaultErr := errors.New("unreachable link in " + linkpath + ":" + link)

	resp, err := http.Get(link)
	if err != nil {
		return defaultErr
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return defaultErr
}
