package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// NOTE: HandlerFunc adalah simple implementation dari interface Handler yang hanya memiliki satu method ServeHTTP

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandlerServeMux(t *testing.T) {
	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	handler.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hi</h1>")
	})

	handler.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Images</h1>")
	})

	handler.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Images Thumbnails</h1>")
	})

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
