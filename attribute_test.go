package htmg

import (
	"bytes"
	"testing"
)

func TestFlagAttr(t *testing.T) {
	buf := new(bytes.Buffer)

	flg := NewFlagAttr("test")

	if flg.IsFlag() == false {
		t.Error("expected it to be flag")
	}

	flg.Write(buf)
	if buf.String() != "test" {
		t.Error("invalid write")
	}
	buf.Reset()

	flg.WriteString(buf)
	if buf.String() != "test" {
		t.Error("invalid write string")
	}
	buf.Reset()

	bad := NewFlagAttr("")

	if err := bad.Write(buf); err == nil {
		t.Error("did not catch invalid name")
	}
	buf.Reset()

	if err := bad.WriteString(buf); err == nil {
		t.Error("did not catch invalid name for string write")
	}
	buf.Reset()
}

func TestAttr(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAttr("foo", "bar")

	if a.IsFlag() == true {
		t.Error("invalid flag")
	}

	a.Write(buf)
	if buf.String() != `foo="bar"` {
		t.Errorf("invalid write %s", buf.String())
	}
	buf.Reset()

	a.WriteString(buf)
	if buf.String() != `foo="bar"` {
		t.Errorf("invalid write string %s", buf.String())
	}
	buf.Reset()

	a = NewAttr("baz", "")
	if a.IsFlag() == false {
		t.Error("expected flag")
	}

	a.Write(buf)
	if buf.String() != `baz` {
		t.Errorf("invalid flag write %s", buf.String())
	}
	buf.Reset()

	a.WriteString(buf)
	if buf.String() != `baz` {
		t.Errorf("invalid flag write string %s", buf.String())
	}
	buf.Reset()
}

func TestSafeAttr(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewSafeAttr("foo", "bar")

	if a.IsFlag() == true {
		t.Error("invalid flag")
	}

	a.Write(buf)
	if buf.String() != `foo="bar"` {
		t.Errorf("invalid write %s", buf.String())
	}
	buf.Reset()

	a.WriteString(buf)
	if buf.String() != `foo="bar"` {
		t.Errorf("invalid write string %s", buf.String())
	}
	buf.Reset()

	a = NewSafeAttr("baz", "")
	if a.IsFlag() == false {
		t.Error("expected flag")
	}

	a.Write(buf)
	if buf.String() != `baz` {
		t.Errorf("invalid flag write %s", buf.String())
	}
	buf.Reset()

	a.WriteString(buf)
	if buf.String() != `baz` {
		t.Errorf("invalid flag write string %s", buf.String())
	}
	buf.Reset()

	a = NewSafeAttr("bad>Name?", "evil \"> value")

	if a.Name() != "badname" {
		t.Errorf("expected clean name, instead %s", a.Name())
	}

	if a.Value() != "evil \"&gt; value" {
		t.Errorf("expected clean value, instead %s", a.Value())
	}

	a.SetValue("value!?<>")
	if a.Value() != "value!?" {
		t.Errorf("expected clean value after change, instead %s", a.Value())
	}
}
