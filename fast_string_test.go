package htmg

import "testing"

func TestStringToBytes(t *testing.T) {
	if stringToBytes("") != nil {
		t.Error("expected nil")
	}

	byt := stringToBytes("foobar")
	if len(byt) != 6 {
		t.Error("invalid size of byte array")
	}
}

func TestBytesToString(t *testing.T) {
	if bytesToString([]byte{}) != "" {
		t.Error("expected empty string")
	} else if bytesToString(nil) != "" {
		t.Error("expected nil to be empty string")
	}

	str := bytesToString([]byte{'f', 'o', 'o'})
	if str != "foo" {
		t.Error("invalid conversion")
	}
}
