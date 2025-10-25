package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

// Recursively convert HTML node to Typst

func htmlToTypst(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	switch n.Data {
	case "b", "strong":
		return "*" + parseChildren(n) + "*"
	case "i", "em":
		return "_" + parseChildren(n) + "_"
	case "p":
		return "" + parseChildren(n) + ""
	// case "", "":
	// 	return "" + parseChildren(n) + ""
	// case "", "":
	// 	return "" + parseChildren(n) + ""
	// case "", "":
	// 	return "" + parseChildren(n) + ""
	// case "", "":
	// 	return "" + parseChildren(n) + ""
	// case "", "":
	// 	return "" + parseChildren(n) + ""
	}
}


func parseChildren(n *html.Node) string {
	//TODO
}
