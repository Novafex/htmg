package htmg

import (
	"bytes"
	"testing"
)

func TestUnsafeText(t *testing.T) {
	buf := new(bytes.Buffer)

	n := NewUnsafeText("foobar")

	n.Write(buf)
	if buf.String() != "foobar" {
		t.Errorf("unexpected results from write %s", buf.String())
	}
	buf.Reset()

	n.WriteString(buf)
	if buf.String() != "foobar" {
		t.Errorf("unexpected results from write string %s", buf.String())
	}
	buf.Reset()

	n.Append(NewUnsafeText(""))
	n.Prepend(NewUnsafeText(""))

	if n.Children() != nil {
		t.Error("expected nil children")
	}
}

func TestSafeText(t *testing.T) {
	buf := new(bytes.Buffer)

	n := NewText("foo<bar>")

	n.Write(buf)
	if buf.String() != "foo" {
		t.Errorf("unexpected results from write %s", buf.String())
	}
	buf.Reset()

	n.WriteString(buf)
	if buf.String() != "foo" {
		t.Errorf("unexpected results from write string %s", buf.String())
	}
	buf.Reset()

	n.Append(NewUnsafeText(""))
	n.Prepend(NewUnsafeText(""))

	if n.Children() != nil {
		t.Error("expected nil children")
	}
}
