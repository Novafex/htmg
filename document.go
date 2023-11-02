package htmg

import (
	"errors"
	"io"
)

var ErrDocRoot = errors.New("no root element in document")

// Doc is short for document and it holds the root node and an optional preamble
// to be displayed at the top of the output.
type Doc struct {
	preamble Writable
	root     *Elem
}

func (d Doc) Write(buf io.Writer) error {
	if d.root == nil {
		return ErrDocRoot
	}

	if d.preamble != nil {
		d.preamble.Write(buf)
	}
	d.root.Write(buf)

	return nil
}

func (d Doc) WriteString(buf io.StringWriter) error {
	if d.root == nil {
		return ErrDocRoot
	}

	if d.preamble != nil {
		d.preamble.WriteString(buf)
	}
	d.root.WriteString(buf)

	return nil
}

func NewDoc(preamble Writable, root *Elem) *Doc {
	return &Doc{
		preamble: preamble,
		root:     root,
	}
}
