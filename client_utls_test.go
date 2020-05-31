package fasthttp

import (
	"testing"
)

func TestGetTLS(t *testing.T) {
	t.Parallel()
	statusCode, _, err := Get(nil, "https://client.tlsfingerprint.io:8443/")
	if err != nil {
		t.Fatal("error was not expected")
	}
	if statusCode != 200 {
		t.Fatalf("StatusCode 200 expected, got %d", statusCode)
	}
}

func TestDoTLS(t *testing.T) {
	t.Parallel()
	req := AcquireRequest()
	resp := AcquireResponse()
	defer func() {
		ReleaseResponse(resp)
		ReleaseRequest(req)
	}()
	req.SetRequestURI("https://client.tlsfingerprint.io:8443/")
	req.Header.SetMethod("GET")
	var err = Do(req, resp)
	if err != nil {
		t.Fatal("error was not expected")
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("StatusCode 200 expected, got %d", resp.StatusCode())
	}
}

func TestClientTLS(t *testing.T) {
	t.Parallel()
	c := &Client{}
	req := AcquireRequest()
	resp := AcquireResponse()
	defer func() {
		ReleaseResponse(resp)
		ReleaseRequest(req)
	}()
	req.SetRequestURI("https://client.tlsfingerprint.io:8443/")
	req.Header.SetMethod("GET")
	err := c.Do(req, resp)
	if err != nil {
		t.Fatal("error was not expected")
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("StatusCode 200 expected, got %d", resp.StatusCode())
	}
}
