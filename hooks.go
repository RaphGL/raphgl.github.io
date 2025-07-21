package main

import (
	"errors"
	"net/http"

	"github.com/gomarkdown/markdown/ast"
)

func CheckLinkIsReachable(filepath string, node ast.Node) error {
	linkNode, ok := node.(*ast.Link)
	if !ok {
		return nil
	}

	link := string(linkNode.Destination)
	defaultErr := errors.New("unreachable link in " + filepath + ":" + link)

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
