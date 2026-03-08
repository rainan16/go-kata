package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// 💡 would  it not be great to be able to test our http handlers?
	// are the response correct?
	// what about the response codes?
	// maybe this can help https://pkg.go.dev/net/http/httptest#example-ResponseRecorder
	t.Run("helloServer", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		helloServer(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", rr.Code)
		}
		if body := rr.Body.String(); body != "Hello World!" {
			t.Fatalf("unexpected body: %q", body)
		}
	})

	t.Run("helloUser normal", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/user/bob/", nil)
		req.SetPathValue("user", "bob")

		helloUser(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", rr.Code)
		}
		if !strings.Contains(rr.Body.String(), "Hello, bob!") {
			t.Fatalf("unexpected body: %q", rr.Body.String())
		}
	})

	t.Run("helloUser teapot", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/user/coffee/", nil)
		req.SetPathValue("user", "coffee")

		helloUser(rr, req)

		if rr.Code != http.StatusTeapot {
			t.Fatalf("expected status 418, got %d", rr.Code)
		}
		if body := rr.Body.String(); body != "I'm a teapot!" {
			t.Fatalf("unexpected body: %q", body)
		}
	})
}
