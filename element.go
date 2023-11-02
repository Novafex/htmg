package htmg

import (
	"fmt"
	"io"
)

// Elem is short for element and is a standard HTML node featuring a tag name
// and optional attributes and children.
//
// It implements both [Writable] and [Node] interfaces
type Elem struct {
	tag      string
	attrs    []Attribute
	children []Node
}

func (e Elem) Write(buf io.Writer) error {
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

	// Do we auto-close this?
	if len(e.children) == 0 {
		// Do self-closing
		buf.Write(STR_SELF_CLOSE)
	} else {
		// End current tag
		buf.Write(CHAR_GT)

		// Write children
		for i, c := range e.children {
			if err = c.Write(buf); err != nil {
				return fmt.Errorf("write child %d in %s element failed: %s", i, e.tag, err.Error())
			}
		}

		// Closing tag
		buf.Write(STR_CLOSE)
		buf.Write(stringToBytes(e.tag))
		buf.Write(CHAR_GT)
	}

	return nil
}

func (e Elem) WriteString(buf io.StringWriter) error {
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

	// Do we auto-close this?
	if len(e.children) == 0 {
		// Do self-closing
		buf.WriteString(" />")
	} else {
		// End current tag
		buf.WriteString(">")

		// Write children
		for i, c := range e.children {
			if err = c.WriteString(buf); err != nil {
				return fmt.Errorf("write child %d in %s element failed: %s", i, e.tag, err.Error())
			}
		}

		// Closing tag
		buf.WriteString("</")
		buf.WriteString(e.tag)
		buf.WriteString(">")
	}

	return nil
}

func (e Elem) Children() []Node {
	return e.children
}

func (e *Elem) Prepend(nodes ...Node) Node {
	e.children = append(nodes, e.children...)
	return e
}

func (e *Elem) Append(nodes ...Node) Node {
	e.children = append(e.children, nodes...)
	return e
}

func NewElem(tag string, attrs ...Attribute) *Elem {
	return &Elem{
		tag:      tag,
		attrs:    attrs,
		children: nil,
	}
}

func NewElemWithChildren(tag string, children ...Node) *Elem {
	return &Elem{
		tag:      tag,
		attrs:    nil,
		children: children,
	}
}
