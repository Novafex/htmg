package htmg

// bad characters not allowed for attribute names.
//
// see https://html.spec.whatwg.org/multipage/syntax.html#attributes-2
var badAttrChars = []byte{'\t', '\n', '\f', ' ', '/', '\\', '>', '=', '=', '\'', '?'}

// sanitizeAttribute removes anything that is not considered a valid HTML/XML
// attribute.
func sanitizeAttribute(str string) string {
	byt := make([]byte, len(str))
	newLen := 0

	for _, b := range stringToBytes(str) {
		if !bytesContain(badAttrChars, b) {
			byt[newLen] = b
			newLen++
		}
	}

	return string(byt[:newLen])
}
