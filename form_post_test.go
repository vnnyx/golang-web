package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	//Bisa juga menggunakan PostFormValue(key) apabila tidak ingin melakukan parsing manual
	//firstName := r.PostFormValue("first_name")

	_, _ = fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	//Membuat request body
	requestBody := strings.NewReader("first_name=Firdaus+Putra&last_name=Kurniyanto")

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)

	//Menambahkan header form post
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
