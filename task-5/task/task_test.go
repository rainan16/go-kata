package task

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleTask_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/task/", nil)
	rr := httptest.NewRecorder()

	HandleTask(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}

func TestHandleTask_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/task/", bytes.NewBufferString("{bad"))
	rr := httptest.NewRecorder()

	HandleTask(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestHandleTask_Valid(t *testing.T) {
	body := bytes.NewBufferString(`{"name":"hello","parallel":2}`)
	req := httptest.NewRequest(http.MethodPost, "/task/", body)
	rr := httptest.NewRecorder()

	HandleTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "result:") {
		t.Fatalf("expected result in response, got %q", rr.Body.String())
	}
}

func TestHandleTask_ParallelDefault(t *testing.T) {
	body := bytes.NewBufferString(`{"name":"hello","parallel":0}`)
	req := httptest.NewRequest(http.MethodPost, "/task/", body)
	rr := httptest.NewRecorder()

	HandleTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "1 parallel") {
		t.Fatalf("expected default parallel=1, got %q", rr.Body.String())
	}
}
