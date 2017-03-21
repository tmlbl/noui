package noui

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testConfig = Config{
	HostName: "localhost",
}

func makeTestRequest(body string) *http.Response {
	app := NewServer()
	dbconnect(testConfig)
	req, _ := http.NewRequest(http.MethodPost, "/api/news",
		bytes.NewBuffer([]byte(body)))
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)
	return res.Result()
}

func expectStatus(t *testing.T, r *http.Response, status int) {
	if r.StatusCode != status {
		t.Errorf("Expected status %d, but got %d", status, r.StatusCode)
	}
}

func TestPostNews(t *testing.T) {
	res := makeTestRequest("hunk")
	expectStatus(t, res, 400)
}
