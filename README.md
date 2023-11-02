# HTMG

Hyper-text Markup Generator is a library for Go that allows fairly efficient
generation of HTML documents for usage with `io.Writer` and strings. Perfect for
dynamic generation of HTML responses to clients over web services.

Features a chained-method style of building "documents" but can be used to make
anything. It is not particular about tag names or attributes so it's up to you
to ensure standards compliance. Most nodes and attributes feature both safe
sanitized versions, and unsafe slightly quicker versions.

## Usage

Install with the usual `go get -u github.com/novafex/htmg` and then start using:

```go
package main

import (
	"os"

	"github.com/novafex/htmg"
)

func main() {
	// Build a tree
	doc := htmg.NewElem("html", htmg.NewAttr("lang", "en")).Append(
		htmg.NewElemWithChildren(
			"body",
			htmg.NewElem("h1").Append(
				htmg.NewText("Hello, World!"),
			),
			htmg.NewText("I'm HTMG from Novafex"),
		),
	)

	doc.Write(os.Stdout)
	// <html lang="en"><body><h1>Hello, World!</h1>I'm HTMG from Novafex</body></html>
}
```

See more running examples in the [examples](./examples/) folder.

## License

MIT License

Copyright (c) 2023 Novafex & Chris Pikul

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
