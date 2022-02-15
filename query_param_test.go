package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		_, _ = fmt.Fprint(w, "Hello")
	} else {
		_, _ = fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestParameterQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Firdaus", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	if firstName == "" && lastName == "" {
		_, _ = fmt.Fprint(w, "Hello")
	} else {
		_, _ = fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
	}
}

func TestMultipleParameterQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?first_name=Firdaus+Putra&last_name=Kurniyanto", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValue(w http.ResponseWriter, r *http.Request) {
	//Untuk mendapatkan multiple parameter value, jangan menggunakan method Get() tapi langsung masukkan key nya
	//karena method Get() hanya mengembalikan value yang pertama saja
	query := r.URL.Query()
	names := query["name"]
	_, _ = fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleParameterValue(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Firdaus&name=Putra&name=Kurniyanto", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
