package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
)

var re = regexp.MustCompile("^\\s*([a-z1-9]*)\\s*color scheme\\s*")

type Color struct {
	scheme   string // x11, svg
	group    string // x11, svg, blue9, blue3
	hexColor string // #ff0011
	name     string // lightblue, 1, 2
}

func main() {
	var colorschemes []Color

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
		id := func(node *html.Node) string { return htmlquery.SelectAttr(node, "id") }
		class := func(node *html.Node) string { return htmlquery.SelectAttr(node, "class") }
		innerText := func(node *html.Node) string { return htmlquery.InnerText(node) }
		title := func(node *html.Node) string { return htmlquery.SelectAttr(node, "title") }

		node := c
		switch scheme := id(node); scheme {
		case "x11":
			{
				for node = node.NextSibling; node != nil; node = node.NextSibling {
					if class(node) == "gv-colors" {
						colors, err := htmlquery.QueryAll(node, "//td")
						if err != nil {
							panic(err)
						}
						for _, colorNode := range colors {
							name := innerText(colorNode)
							hexColor := title(colorNode)
							colorschemes = append(colorschemes, Color{
								scheme:   scheme,
								group:    scheme,
								hexColor: hexColor,
								name:     name,
							})
						}

						break
					}
				}
			}
		case "svg":
			{
				for node = node.NextSibling; node != nil; node = node.NextSibling {
					if class(node) == "gv-colors" {
						colors, err := htmlquery.QueryAll(node, "//td")
						if err != nil {
							panic(err)
						}
						for _, colorNode := range colors {
							name := innerText(colorNode)
							hexColor := title(colorNode)
							colorschemes = append(colorschemes, Color{
								scheme:   scheme,
								group:    scheme,
								hexColor: hexColor,
								name:     name,
							})
						}

						break
					}
				}
			}
		case "brewer":
			for node = node.NextSibling; node != nil; node = node.NextSibling {
				if node.Type == html.TextNode {
					text := node.Data
					// text must be in form like: \n\nylorrd5 color scheme
					list := re.FindStringSubmatch(text)
					var group string
					if len(list) >= 2 {
						group = list[1]
						// find colors
						for node = node.NextSibling; node != nil; node = node.NextSibling {
							if class(node) == "gv-colors" {
								colors, err := htmlquery.QueryAll(node, "//td")
								if err != nil {
									panic(err)
								}
								for _, colorNode := range colors {
									name := innerText(colorNode)
									hexColor := title(colorNode)
									colorschemes = append(colorschemes, Color{
										scheme:   scheme,
										group:    group,
										hexColor: hexColor,
										name:     name,
									})
								}
								break
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(colorschemes)
}
