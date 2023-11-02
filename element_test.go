package htmg

import (
	"bytes"
	"testing"
)

func TestElemWrite(t *testing.T) {
	buf := new(bytes.Buffer)

	e := NewElem("foo", NewAttr("bar", "baz"))

	// With children and attrs
	e.Append(NewUnsafeText("fiz"))
	e.Write(buf)
	if buf.String() != `<foo bar="baz">fiz</foo>` {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()
}

func TestElemWriteString(t *testing.T) {
	buf := new(bytes.Buffer)

	e := NewElem("foo", NewAttr("bar", "baz"))

	// With children and attrs
	e.Prepend(NewUnsafeText("fiz"))
	e.WriteString(buf)
	if buf.String() != `<foo bar="baz">fiz</foo>` {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()
}

func TestNewElemWithChildren(t *testing.T) {
	buf := new(bytes.Buffer)

	e := NewElemWithChildren("foo", NewUnsafeText("fiz"))

	if len(e.Children()) != 1 {
		t.Error("expected children")
	}

	e.WriteString(buf)
	if buf.String() != `<foo>fiz</foo>` {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()
}
