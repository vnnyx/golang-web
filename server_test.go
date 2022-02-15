package golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	//Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
