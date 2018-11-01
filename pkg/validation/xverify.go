package validation

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var (
	// Default XVerify endpoint
	endpoint = "http://www.xverify.com/services/emails/verify"

	// ErrInvalidResponseCode is returned when response code is different than 200
	ErrInvalidResponseCode = errors.New("Invalid response code")
)

// XVerifyResponse structure returned in the HTTP call
type XVerifyResponse struct {
	Email struct {
		Error   int    `json:"error"`
		Message string `json:"message"`
	} `json:"email"`
}

type XVerify struct {
	apiKey  string
	domain  string
	timeout time.Duration
}

// NewXVerify creates an XVerify validator
func NewXVerify(apiKey, domain string, timeout time.Duration) *XVerify {
	return &XVerify{
		apiKey:  apiKey,
		domain:  domain,
		timeout: timeout,
	}
}

// SetEndpoint overrides the default endpoint
func (x *XVerify) SetEndpoint(e string) {
	endpoint = e
}

// Validate the email against XVerify
func (x *XVerify) Validate(email string) (bool, error) {
	response, err := x.verify(email)
	if err != nil {
		return false, err
	}

	return response.Email.Error == 0, nil
}

// Verify call to XVerify services
func (x *XVerify) verify(email string) (*XVerifyResponse, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("type", "json")
	q.Add("email", email)
	q.Add("apikey", x.apiKey)
	q.Add("domain", x.domain)
	req.URL.RawQuery = q.Encode()

	client := http.Client{Timeout: x.timeout}
	response, err := client.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, ErrInvalidResponseCode
	}

	var resp XVerifyResponse
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
