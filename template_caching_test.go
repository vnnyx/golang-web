package golang_web

import (
	"embed"
	"fmt"
	//Package untuk melakukan escape tidak boleh "text/template"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templatesFile embed.FS

var myTemplates = template.Must(template.ParseFS(templatesFile, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	_ = myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Templates")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
