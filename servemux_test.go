package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	//Membuat beberapa endpoint menggunakan servemux
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello, World")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hi")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/image/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Image")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/image/thumbnails", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Thumbnail")
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, request.Method)
		_, _ = fmt.Fprint(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
