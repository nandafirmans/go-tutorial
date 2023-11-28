package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./12_templates/if.html"))
	t.ExecuteTemplate(w, "if.html", Page{
		Title: "Lorem Ipsum",
		Name:  "Nanda Firmansyah",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionIfCompare(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./12_templates/if-compare.html"))
	t.ExecuteTemplate(w, "if-compare.html", map[string]interface{}{
		"Title": "Lorem Ipsum",
		"Value": 80,
	})
}

func TestTemplateActionIfCompare(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIfCompare(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./12_templates/range.html"))
	t.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Title": "Lorem Ipsum",
		"Names": []string{"Agus", "Budi", "Cecep", "Doni", "Eko"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
