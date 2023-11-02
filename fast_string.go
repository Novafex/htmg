package htmg

import "unsafe"

// stringToBytes provides a fast, but dangerous, way to convert strings into
// byte-slices.
func stringToBytes(str string) []byte {
	if str == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

// bytesToString provides a fast, but dangerous, way to convert byte-slices into
// strings.
func bytesToString(byt []byte) string {
	if len(byt) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(byt), len(byt))
}
