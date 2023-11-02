package htmg

import "io"

// Writable is an interface for any type within [htmg] that can be written to an
// [io] compatible interface. It includes both the [io.Writer] interface and the
// [io.StringWriter] interface for dealing with bytes vs strings. A lot of htmg
// is string based so the [WriteString] interface should be preferred.
//
// Both implementations should result in the same results, just by different
// type means. That is, the [Write] and [WriteString] should produce the same
// outcomes using their specific types of []byte and string.
type Writable interface {
	// Write accepts an [io.Writer] interface for writing byte slices to the
	// output buffer. It should return an error if one occurs which will halt
	// the encoding process.
	Write(buf io.Writer) error

	// WriteString accepts an [io.StringWriter] interface for writing strings to
	// the output buffer. It should return an error if one occurs which will
	// half the encoding process.
	WriteString(buf io.StringWriter) error
}

// WritableString is a helper type for pure text content that can be written out
// to the buffers.
type WritableString struct {
	content string
}

func (w WritableString) Write(buf io.Writer) error {
	buf.Write(stringToBytes(w.content))
	return nil
}

func (w WritableString) WriteString(buf io.StringWriter) error {
	buf.WriteString(w.content)
	return nil
}

// NewString contructs a new [WritableString] using the contents provided.
//
// [WritableString] is just a helper type that can write text directly as a
// [Writable] interfaced object.
func NewString(text string) Writable {
	return &WritableString{text}
}
