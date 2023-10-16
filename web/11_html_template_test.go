package web

import (
	"embed"
	_ "embed"
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
	t := template.Must(template.ParseFiles("./11_templates/simple.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello Template")
}

func TestHtmlTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./11_templates/*.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed 11_templates/*.html
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "11_templates/*.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateData(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./11_templates/user.html"))
	t.ExecuteTemplate(w, "user.html", map[string]interface{}{
		"Title": "Lorem Ipsum",
		"Name":  "Nanda",
		"Address": map[string]interface{}{
			"Street": "Jalan jalan jalan",
		},
	})
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateData(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./11_templates/user.html"))
	t.ExecuteTemplate(w, "user.html", Page{
		Title: "Lorem Ipsum",
		Name:  "Nanda Firmansyah",
		Address: Address{
			Street: "Jalan jalan jalan",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8087", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
