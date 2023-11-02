package htmg

// bytesContain returns true if the slice contains the given byte
func bytesContain(slice []byte, byt byte) bool {
	for i := range slice {
		if slice[i] == byt {
			return true
		}
	}
	return false
}

// bytesContainRune returns true if the slice contains the given rune
func bytesContainRune(slice []byte, r rune) bool {
	byt := byte(r)
	for i := range slice {
		if slice[i] == byt {
			return true
		}
	}
	return false
}
