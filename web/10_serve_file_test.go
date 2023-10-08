package web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	hasName := r.URL.Query().Get("name") != ""

	if hasName {
		http.ServeFile(w, r, "./10_resources/ok.html")
	} else {
		http.ServeFile(w, r, "./10_resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8087",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed 10_resources/ok.html
var resOk string

//go:embed 10_resources/notfound.html
var resNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	hasName := r.URL.Query().Get("name") != ""

	if hasName {
		fmt.Fprint(w, resOk)
	} else {
		fmt.Fprint(w, resNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8087",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
