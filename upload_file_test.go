package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	//default memory adalah 32 MB
	//Namun bisa diganti dengan
	//request.ParseMultipartForm(100 << 20) -> 100 MB

	//Mengambil file
	//"file" disamakan dengan fieldname yang ada di .gohtml
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	//Menyimpan file kedalam sebuah directory
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	//Memindahkan file upload kedalah fileDestination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	//Mengambil yang bukan file
	//"name" disesuaikan dengan fieldname pada file .gohtml
	name := request.PostFormValue("name")
	_ = myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//Unit test
//go:embed resources/test.jpg
var uploadFileTest []byte

func TestUpload(t *testing.T) {
	//Membuat body yang akan dikirim
	body := new(bytes.Buffer)

	//Memasukkan sesuatu kedalan body
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("name", "Firdaus Putra Kurniyanto")
	file, _ := writer.CreateFormFile("file", "FILE.png")
	_, _ = file.Write(uploadFileTest)
	_ = writer.Close()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", body)
	//Set header Content-Type
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
