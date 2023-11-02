package main

import (
	"os"

	"github.com/novafex/htmg"
)

func main() {
	// Build a tree
	doc := htmg.NewDoc(
		htmg.NewString("<!DOCTYPE html>"),
		htmg.NewElem("html", htmg.NewAttr("lang", "en")).Append(
			htmg.NewElemWithChildren(
				"body",
				htmg.NewElem("h1").Append(
					htmg.NewText("Hello, World!"),
				),
				htmg.NewText("I'm HTMG from Novafex"),
			),
		).(*htmg.Elem),
	)

	doc.Write(os.Stdout)
	// <!DOCTYPE html><html lang="en"><body><h1>Hello, World!</h1>I'm HTMG from Novafex</body></html>
}
