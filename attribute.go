package htmg

import (
	"errors"
	"io"

	"github.com/sym01/htmlsanitizer"
)

var ErrAttrName = errors.New("attribute is missing a name")

type Attribute interface {
	Write(buf io.Writer) error
	WriteString(buf io.StringWriter) error

	IsFlag() bool
}

// FlagAttr is short for boolean attribute and it is a named attribute that
// HTML should accept as a boolean flag declaring the attribute "true" if
// present.
type FlagAttr struct {
	Name string
}

func (a FlagAttr) Write(buf io.Writer) error {
	if len(a.Name) == 0 {
		return ErrAttrName
	}
	buf.Write(stringToBytes(a.Name))
	return nil
}

func (a FlagAttr) WriteString(buf io.StringWriter) error {
	if len(a.Name) == 0 {
		return ErrAttrName
	}
	buf.WriteString(a.Name)
	return nil
}

func (a FlagAttr) IsFlag() bool {
	return true
}

func NewFlagAttr(name string) *FlagAttr {
	return &FlagAttr{name}
}

// Attr is short for attribute and is a key-value pair used by HTML/XML to set
// attributes of a node. They are encoded as `name="value"`. Note that if no
// value (or empty) is used then it will be treated as a flag attribute.
type Attr struct {
	Name  string
	Value string
}

func (a Attr) Write(buf io.Writer) error {
	if len(a.Name) == 0 {
		return ErrAttrName
	}
	buf.Write(stringToBytes(a.Name))

	if len(a.Value) > 0 {
		buf.Write(STR_EQ_QT)
		buf.Write(stringToBytes(a.Value))
		buf.Write(CHAR_EQ)
	}
	return nil
}

func (a Attr) WriteString(buf io.StringWriter) error {
	if len(a.Name) == 0 {
		return ErrAttrName
	}
	buf.WriteString(a.Name)

	if len(a.Value) > 0 {
		buf.WriteString(`="`)
		buf.WriteString(a.Value)
		buf.WriteString(`"`)
	}
	return nil
}

func (a Attr) IsFlag() bool {
	return len(a.Value) > 0
}

func NewAttr(name, value string) *Attr {
	return &Attr{
		Name:  name,
		Value: value,
	}
}

// SafeAttr is an attribute that is considered "safe". It is immutable and has
// it's name and values HTML sanitized when created.
//
// Use [NewSafeAttr] to create them
type SafeAttr struct {
	// boolean declares if this attribute has a value or should be treated as a
	// flag instead.
	bool

	name  string
	value string
}

func (a SafeAttr) Write(buf io.Writer) error {
	if len(a.name) == 0 {
		return ErrAttrName
	}
	buf.Write(stringToBytes(a.name))
	if a.bool {
		buf.Write(STR_EQ_QT)
		buf.Write(stringToBytes(a.value))
		buf.Write(CHAR_EQ)
	}
	return nil
}

func (a SafeAttr) WriteString(buf io.StringWriter) error {
	if len(a.name) == 0 {
		return ErrAttrName
	}
	buf.WriteString(a.name)
	if a.bool {
		buf.WriteString(`="`)
		buf.WriteString(a.value)
		buf.WriteString(`"`)
	}
	return nil
}

func (a SafeAttr) IsFlag() bool {
	return len(a.value) > 0
}

// Name returns the current name of the attribute
func (a SafeAttr) Name() string {
	return a.name
}

// Value returns the current value of the attribute
func (a SafeAttr) Value() string {
	return a.value
}

// SetValue sets the value of the attribute and does so by sanitizing the input
// for HTML. Returns an error from [htmlsanitizer] if it encounters one.
func (a *SafeAttr) SetValue(val string) error {
	var err error
	a.name, err = htmlsanitizer.SanitizeString(val)
	return err
}

// NewSafeAttr creates a new SafeAttr attribute that has it's name and value
// sanitized. The value is optional and if an empty string is used, the attribute
// will be treated as a boolean flag.
func NewSafeAttr(name, value string) *SafeAttr {
	attr := &SafeAttr{
		bool: len(value) > 0,
		name: sanitizeAttribute(name),
	}

	if len(value) > 0 {
		safeValue, err := htmlsanitizer.SanitizeString(value)
		if err != nil {
			panic(err)
		}
		attr.value = safeValue
	}

	return attr
}
