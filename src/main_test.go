package main

import (
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	req := httptest.NewRequest("GET", "http://localhost:3000/hello", nil)
	w := httptest.NewRecorder()
	hello(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Error("response is not OK")
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		t.Error("parsing error")
		return
	}

	text := doc.Find("body").Text()
	if text != "Hello World!" {
		t.Error("unexpected data")
		return
	}
}
