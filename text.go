package htmg

import (
	"io"

	"github.com/sym01/htmlsanitizer"
)

// Text is a [Node] type that writes text to the output buffer like any other
// [Writable] does. The difference is it's immutable and is sanitized when
// created via [NewText]. Note that it "passes" as a Node but does not behave
// like one. This is to trick the encoders into using it anyways.
type Text struct {
	string
}

func (n Text) Write(buf io.Writer) error {
	buf.Write(stringToBytes(n.string))
	return nil
}

func (n Text) WriteString(buf io.StringWriter) error {
	buf.WriteString(n.string)
	return nil
}

func (n Text) Children() []Node {
	return nil
}

func (n *Text) Prepend(nodes ...Node) Node {
	return n
}

func (n *Text) Append(nodes ...Node) Node {
	return n
}

// NewUnsafeText returns a new Text object with the contents copied in. It is
// unsafe in that the content is not sterilized. It is recommended to use
// [NewText] instead if possible.
func NewUnsafeText(content string) *Text {
	return &Text{content}
}

// NewText returns a new Text object with the contents sanitized and copied in.
//
// It might be slightly slower than [NewUnsafeText] but it's safer.
func NewText(content string) *Text {
	safe, err := htmlsanitizer.SanitizeString(content)
	if err != nil {
		panic(err)
	}

	return &Text{safe}
}
