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

func TestPostNews(t *testing.T) {
	app := NewServer()
	dbconnect(testConfig)
	req, _ := http.NewRequest(http.MethodPost, "/api/news",
		bytes.NewBuffer([]byte("hunnunun")))
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)
}
