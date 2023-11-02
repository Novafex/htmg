package html

import "github.com/novafex/htmg"

var DocPreamble = htmg.NewString("<!DOCTYPE html>")

// Doc creates a new HTML document with a preamble and root <html /> node applied.
func Doc(children ...htmg.Node) *htmg.Doc {
	return htmg.NewDoc(DocPreamble, htmg.NewElemWithChildren("html", children...))
}
