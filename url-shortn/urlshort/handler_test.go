package urlshort

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var pathToUrls = map[string]string{
	"/google":      "https://www.google.com.br",
	"/facebook":    "https://www.facebook.com.br",
	"/is-it-virus": "https://www.virustotal.com",
}

func TestMapHandler(t *testing.T) {
	path := "/google"

	handler := MapHandler(pathToUrls, http.NewServeMux())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)
	handler(w, r)

	res := w.Result()

	redirectedTo := res.Header.Get("Location")
	expectedRedirect := pathToUrls[path]

	if redirectedTo != expectedRedirect {
		t.Errorf("expected redirect to %v but got %s", expectedRedirect, redirectedTo)
	}
}

func TestMapHandlerCallback(t *testing.T) {
	path := "/unknown"

	expectedBody := "fallback called!"
	fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedBody))
	})

	handler := MapHandler(pathToUrls, fallback)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)
	handler(w, r)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Error(err)
	}

	pBody := string(body)
	if pBody != expectedBody {
		t.Errorf("expected body result to be %s but got %s", expectedBody, pBody)
	}
}

func TestYamlHandler(t *testing.T) {
	path := "/mine"
	expectedRedirect := "https://www.minecraft.net"

	ymlBytes := []byte(fmt.Sprintf(`- path: %s
  url: %s
- path: /gtav
  url: https://www.rockstargames.com/gta-v`, path, expectedRedirect))

	handler, err := YAMLHandler(ymlBytes, http.NewServeMux())
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)
	handler(w, r)

	redirectedTo := w.Result().Header.Get("Location")

	if redirectedTo != expectedRedirect {
		t.Errorf("expected redirect to be %s but got %s", expectedRedirect, redirectedTo)
	}
}

func TestYamlHandlerFallback(t *testing.T) {
	ymlBytes := []byte(`- path: /gtav
  url: https://www.rockstargames.com/gta-v`)

	fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello-world", http.StatusSeeOther)
	})

	handler, err := YAMLHandler(ymlBytes, fallback)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	handler(w, r)

	redirectedTo := w.Result().Header.Get("Location")
	expectedRedirect := "/hello-world"

	if redirectedTo != expectedRedirect {
		t.Errorf("expected redirect to be %s but got %s", expectedRedirect, redirectedTo)
	}
}

func TestJSONHandler(t *testing.T) {
	path := "/mine"
	expectedRedirect := "https://www.minecraft.net"

	jsonBytes := []byte(fmt.Sprintf(`[
		{ "path": "%s", "url": "%s" },
		{ "path": "/gtav", "url": "https://www.rockstargames.com/gta-v" }
	]`, path, expectedRedirect))

	handler, err := JSONHandler(jsonBytes, http.NewServeMux())
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)
	handler(w, r)

	redirectedTo := w.Result().Header.Get("Location")

	if redirectedTo != expectedRedirect {
		t.Errorf("expected redirect to be %s but got %s", expectedRedirect, redirectedTo)
	}
}

func TestJSONHandlerFallback(t *testing.T) {
	ymlBytes := []byte(`[{ "path": "/gtav", "url": "https://www.rockstargames.com/gta-v" }]`)

	fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello-world", http.StatusSeeOther)
	})

	handler, err := JSONHandler(ymlBytes, fallback)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	handler(w, r)

	redirectedTo := w.Result().Header.Get("Location")
	expectedRedirect := "/hello-world"

	if redirectedTo != expectedRedirect {
		t.Errorf("expected redirect to be %s but got %s", expectedRedirect, redirectedTo)
	}
}
