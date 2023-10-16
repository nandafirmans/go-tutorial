package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	simpleTemplate := template.New("SIMPLE")

	// t, err := simpleTemplate.Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	// NOTE: lebih simple ga perlu cek error karena sudah dihandle didalam fungsi template.Must
	t := template.Must(simpleTemplate.Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello World")
}

func TestHtmlTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./11_templates/simple.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.html", "Hello Template")
}

func TestHtmlTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
