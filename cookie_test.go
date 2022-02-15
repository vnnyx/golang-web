package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "X-Pama",
		Value: r.URL.Query().Get("name"),
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	_, _ = fmt.Fprint(w, "Success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Pama")
	if err != nil {
		_, _ = fmt.Fprint(w, "No cookie")
	} else {
		_, _ = fmt.Fprintf(w, "Cookie name: %s \nCookie Value: %s", cookie.Name, cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/8080/?name=firdaus%20putra", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookie := recorder.Result().Cookies()

	for _, cookies := range cookie {
		fmt.Printf("Cookie Name: %s \nCookie Value: %s\n", cookies.Name, cookies.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Pama"
	cookie.Value = "Firdaus Putra Kurniyanto"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
