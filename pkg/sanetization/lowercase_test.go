package sanetization

import "testing"

func TestLowercaseSanity(t *testing.T) {
	sanetizer := &Lowercase{}
	sanetizedEmail, _ := sanetizer.Sanitize("eMail@email.coM")
	if sanetizedEmail != "email@email.com" {
		t.Errorf("Invalid basic sanity: %s", sanetizedEmail)
	}
}
