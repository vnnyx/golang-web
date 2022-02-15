package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello World")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	//Mendapatkan result
	response := recorder.Result()
	//Membaca result
	body, _ := io.ReadAll(response.Body)
	//Karena body merupakan byte dan kita hanya perlu string, maka ubah ke string
	bodyString := string(body)

	fmt.Println(bodyString)

}
