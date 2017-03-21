package noui

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostNews(t *testing.T) {
	app := NewServer()
	Serve()
	req, _ := http.NewRequest(http.MethodPost, "/api/news",
		bytes.NewBuffer([]byte("hunnunun")))
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)
}
