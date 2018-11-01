package validation

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestItReturnsAnErrorIfCanNotCreateTheRequest(t *testing.T) {
	v := NewXVerify("apiKey", "domain", 1*time.Second)
	v.SetEndpoint("invalid endpoint")
	_, err := v.Validate("email@email.com")
	if err == nil {
		t.Error("An error should be returned if can not create the request")
	}
}

func TestItReturnsAnErrorIfTheRequestFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}))
	defer ts.Close()

	v := NewXVerify("apiKey", "domain", 1*time.Second)
	v.SetEndpoint(ts.URL)

	_, err := v.Validate("email@email.com")
	if err != ErrInvalidResponseCode {
		t.Error("A 'ErrInvalidResponseCode' error should be returned if the response code is different than 200")
	}
}

func TestItReturnsAnErrorIfCanNotDecodeTheResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()

	v := NewXVerify("apiKey", "domain", 1*time.Second)
	v.SetEndpoint(ts.URL)

	_, err := v.Validate("email@email.com")
	if err == nil {
		t.Error("A error should be returned if can't decode the response")
	}
}

func TestItReturnsAnErrorIfExceedsTheTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"email\":{\"error\":0, \"message\":\"valid email\"}}"))
		time.Sleep(100 * time.Millisecond)
	}))
	defer ts.Close()

	v := NewXVerify("apiKey", "domain", 1*time.Millisecond)
	v.SetEndpoint(ts.URL)

	_, err := v.Validate("email@email.com")
	if err == nil {
		t.Error("An error should be returned if the request exceeds the timeout")
	}
}

func TestItReturnsXVerifyResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"email\":{\"error\":0, \"message\":\"valid email\"}}"))
	}))
	defer ts.Close()

	v := NewXVerify("apiKey", "domain", 1*time.Second)
	v.SetEndpoint(ts.URL)

	valid, err := v.Validate("email@email.com")
	if err != nil {
		t.Error("No error should be returned")
	}
	if !valid {
		t.Error("The result should be valid if there are no errors in xverify response")
	}
}
