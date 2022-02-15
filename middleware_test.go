package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Middleware")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Middleware")
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(writer, "Error %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Execute")
		_, _ = fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Execute")
		_, _ = fmt.Fprint(writer, "Hello Foo")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Execute")
		panic("Ups")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errorHandler := new(ErrorHandler)
	errorHandler.Handler = logMiddleware

	server := new(http.Server)
	server.Addr = "localhost:8080"
	server.Handler = errorHandler

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
