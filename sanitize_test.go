package htmg

import "testing"

func TestSanitizeAttribute(t *testing.T) {
	check := "bad?attr/name"

	if sanitizeAttribute(check) != "badattrname" {
		t.Error("invalid returned value")
	}

	if sanitizeAttribute(string(badAttrChars)) != "" {
		t.Error("did not remove all bad characters")
	}
}
