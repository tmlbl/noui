package noui

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testConfig() Config {
	return Config{
		HostName: "localhost",
	}
}

func makeTestRequest(method, path, body string) *http.Request {
	c := testConfig()
	c.DBName = "noui"
	dbconnect(c)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	return req
}

func execRequest(req *http.Request) *http.Response {
	app := NewServer()
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)
	return res.Result()
}

func expectStatus(t *testing.T, r *http.Response, status int) {
	if r.StatusCode != status {
		t.Errorf("Expected status %d, but got %d", status, r.StatusCode)
	}
}

func TestBadJSON(t *testing.T) {
	req := makeTestRequest("POST", "/news", "blablabla")
	res := execRequest(req)
	expectStatus(t, res, 400)
}

func TestPostNews(t *testing.T) {
	news := News{
		Model: Model{
			Namespace: "test_news",
		},
		Headline: "A test news item",
		Content:  "u already know",
	}
	js, _ := json.Marshal(news)
	req := makeTestRequest("POST", "/news", string(js))
	res := execRequest(req)
	expectStatus(t, res, 200)

	// Fetch the news collection
	req = makeTestRequest("GET", "/news/test_news", "")
	res = execRequest(req)
}
