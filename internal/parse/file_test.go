package parse

import (
	"testing"
)

func TestItReturnsAnErrorIfCanNotOpenTheFile(t *testing.T) {
	_, err := From("invalid_file")
	if err == nil {
		t.Error("An error should be returned if the file does not exist")
	}
}

func TestItReturnsAllTheLines(t *testing.T) {
	lines, err := From("./file_test.txt")
	if err != nil {
		t.Error("No error should be returned")
	}
	if len(lines) != 4 {
		t.Errorf("Expected 4 lines, given: %d", len(lines))
	}
}

func TestItReturnsAnErrorIfCanNotOpenTheRemoteFile(t *testing.T) {
	_, err := From("http://invalid_file.com/file")
	if err == nil {
		t.Error("An error should be returned if URL file does not exist")
	}
}

func TestItReturnsAllTheLinesFromARemoteFile(t *testing.T) {
	_, err := From("https://raw.githubusercontent.com/martenson/disposable-email-domains/master/allowlist.conf")
	if err != nil {
		t.Error("No error should be returned")
	}
}
