package sanetization

import "testing"

func TestTrimSanity(t *testing.T) {
	sanitizer := &Trim{}
	sanetizedEmail, _ := sanitizer.Sanitize("   email@email.com ")
	if sanetizedEmail != "email@email.com" {
		t.Errorf("Invalid trim sanity: %s", sanetizedEmail)
	}
}
