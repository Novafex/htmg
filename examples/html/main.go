package main

import (
	"os"

	"github.com/novafex/htmg"
	"github.com/novafex/htmg/html"
)

func main() {
	doc := html.Doc(htmg.NewText("Hello, World!"))

	doc.Write(os.Stdout)
}
