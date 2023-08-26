package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprint(w, "Hello ", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3001/hello?name=Afif", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func SayHelloFullName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	firstName := query.Get("first_name")
	lastName := query.Get("last_name")

	if firstName == "" && lastName == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprint(w, "Hello ", firstName, " ", lastName)
	}
}

func TestQueryParam2(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:3001/hello?first_name=Afif&last_name=AlHanif",
		nil,
	)
	recorder := httptest.NewRecorder()

	SayHelloFullName(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, ", "))
}

func TestQueryParamMultiple(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:3001/hello?name=Muhammad&name=Afif&name=AlHanif",
		nil,
	)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
