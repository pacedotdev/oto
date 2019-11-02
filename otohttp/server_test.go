package otohttp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/matryer/is"
)

func TestServer(t *testing.T) {
	is := is.New(t)
	srv := NewServer()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"greeting":"Hi Mat"}`))
	})
	srv.Register("Service", "Method", h)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/oto/Service.Method", strings.NewReader(`{"name":"Mat"}`))
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `{"greeting":"Hi Mat"}`)
}

func TestEncode(t *testing.T) {
	is := is.New(t)
	data := struct {
		Greeting string `json:"greeting"`
	}{
		Greeting: "Hi there",
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/oto/Service.Method", strings.NewReader(`{"name":"Mat"}`))
	err := Encode(w, r, http.StatusOK, data)
	is.NoErr(err)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `{"greeting":"Hi there"}`)
	is.Equal(w.HeaderMap.Get("Content-Type"), "application/json; chatset=utf-8")
}

func TestDecode(t *testing.T) {
	is := is.New(t)
	type r struct {
		Name string
	}
	j := `[
		{"name": "Mat"},
		{"name": "David"},
		{"name": "Aaron"}
	]`
	req, err := http.NewRequest(http.MethodPost, "/service/method", strings.NewReader(j))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/json")
	var requestObjects []r
	err = remotohttp.Decode(req, &requestObjects)
	is.NoErr(err)
	is.Equal(len(requestObjects), 3)
	is.Equal(requestObjects[0].Name, "Mat")
	is.Equal(requestObjects[1].Name, "David")
	is.Equal(requestObjects[2].Name, "Aaron")
}
