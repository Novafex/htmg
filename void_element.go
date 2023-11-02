package htmg

import (
	"fmt"
	"io"
)

// Void is a special type of element that features a tag, and attributes, but
// can never contain children.
//
// It implements both [Writable] and [Node] interfaces
type Void struct {
	tag   string
	attrs []Attribute
}

func (e Void) Write(buf io.Writer) error {
	// Start opening tag
	buf.Write(CHAR_LT)
	buf.Write(stringToBytes(e.tag))

	// Write attributes
	var err error
	for i, a := range e.attrs {
		buf.Write(CHAR_SPACE)
		if err = a.Write(buf); err != nil {
			return fmt.Errorf("write attribute %d in %s element failed: %s", i, e.tag, err.Error())
		}
	}

	// End current tag
	buf.Write(STR_SELF_CLOSE)

	return nil
}

func (e Void) WriteString(buf io.StringWriter) error {
	// Start opening tag
	buf.WriteString("<")
	buf.WriteString(e.tag)

	// Write attributes
	var err error
	for i, a := range e.attrs {
		buf.WriteString(" ")
		if err = a.WriteString(buf); err != nil {
			return fmt.Errorf("write attribute %d in %s element failed: %s", i, e.tag, err.Error())
		}
	}

	buf.WriteString(" />")

	return nil
}

func (e Void) Children() []Node {
	return nil
}

func (e *Void) Prepend(nodes ...Node) Node {
	panic("void elements can never contain children")
}

func (e *Void) Append(nodes ...Node) Node {
	panic("void elements can never contain children")
}

func NewVoid(tag string, attrs ...Attribute) *Void {
	return &Void{
		tag:   tag,
		attrs: attrs,
	}
}
