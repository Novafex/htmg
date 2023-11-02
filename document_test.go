package htmg

import (
	"bytes"
	"testing"
)

func TestDocWrite(t *testing.T) {
	buf := new(bytes.Buffer)

	d := NewDoc(nil, NewElem("foo"))
	d.Write(buf)
	if buf.String() != "<foo></foo>" {
		t.Errorf("invalid write response '%s'", buf.String())
	}
	buf.Reset()

	d = NewDoc(NewString("<!DOCTYPE html>"), NewElem("foo"))
	d.Write(buf)
	if buf.String() != "<!DOCTYPE html><foo></foo>" {
		t.Errorf("invalid write response '%s'", buf.String())
	}
	buf.Reset()
}

func TestDocWriteString(t *testing.T) {
	buf := new(bytes.Buffer)

	d := NewDoc(nil, NewElem("foo"))
	d.WriteString(buf)
	if buf.String() != "<foo></foo>" {
		t.Errorf("invalid write response '%s'", buf.String())
	}
	buf.Reset()

	d = NewDoc(NewString("<!DOCTYPE html>"), NewElem("foo"))
	d.WriteString(buf)
	if buf.String() != "<!DOCTYPE html><foo></foo>" {
		t.Errorf("invalid write response '%s'", buf.String())
	}
	buf.Reset()
}
