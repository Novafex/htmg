package htmg

import "testing"

func TestBytesContain(t *testing.T) {
	tst := []byte("abc123")

	if !bytesContain(tst, '1') {
		t.Error("expected it to find it")
	}

	if bytesContain(tst, 'z') {
		t.Error("expected it to not find it")
	}
}

func TestBytesContainRune(t *testing.T) {
	tst := []byte("abc💼")

	if !bytesContainRune(tst, '💼') {
		t.Error("expected it to find it")
	}

	if bytesContainRune(tst, 'z') {
		t.Error("expected it to not find it")
	}
}
