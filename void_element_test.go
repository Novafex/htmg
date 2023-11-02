package htmg

import (
	"bytes"
	"testing"
)

func TestVoidWrite(t *testing.T) {
	buf := new(bytes.Buffer)

	e := NewVoid("foo")

	// Self closing no attrs
	e.Write(buf)
	if buf.String() != "<foo />" {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()

	// Self closing with attrs
	e = NewVoid("foo", NewAttr("bar", "baz"))
	e.Write(buf)
	if buf.String() != `<foo bar="baz" />` {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()
}

func TestVoidWriteString(t *testing.T) {
	buf := new(bytes.Buffer)

	e := NewVoid("foo")

	// Self closing no attrs
	e.WriteString(buf)
	if buf.String() != "<foo />" {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()

	// Self closing with attrs
	e = NewVoid("foo", NewAttr("bar", "baz"))
	e.WriteString(buf)
	if buf.String() != `<foo bar="baz" />` {
		t.Errorf("invalid write '%s'", buf.String())
	}
	buf.Reset()
}
