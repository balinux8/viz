package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
)

func main() {
	resp, err := http.Get("https://graphviz.org/doc/info/colors.html")
	if err != nil {
		panic(err)
	}
	root, err := htmlquery.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	colorSchemes, err := htmlquery.QueryAll(root, "//h2")
	if err != nil {
		panic(err)
	}
	if colorSchemes == nil {
		return
	}

	for _, c := range colorSchemes {
		fmt.Println(htmlquery.SelectAttr(c, "id"))
	}
}
