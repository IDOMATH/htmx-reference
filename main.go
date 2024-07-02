package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/idomath/htmx-reference/middleware"
)

type Tmpl struct {
	tmpl *template.Template
}

func (t *Tmpl) Render(w io.Writer, name string, data interface{}) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func newTemplate() *Tmpl {
	return &Tmpl{tmpl: template.Must(template.ParseGlob("*.html"))}
}

type TemplateData struct {
	Count int
}

var data TemplateData

var templates map[string]*template.Template

func main() {

	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(router),
	}

	templates = make(map[string]*template.Template)
	data.Count = 0

	tmpl, err := template.New("index").ParseFiles("./views/index.html")
	if err != nil {
		panic("could not make index template")
	}
	templates["index"] = tmpl

	router.HandleFunc("GET /", handleHome)
	router.HandleFunc("POST /count", handlePostCount)

	fmt.Println("server started on port 8080")
	log.Fatal(server.ListenAndServe())
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)

	err := templates["index"].Execute(w, data)

	if err != nil {
		log.Fatal(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return
	}

}

func handlePostCount(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	data.Count++

	err := templates["index"].Execute(w, data)

	if err != nil {
		log.Fatal(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return
	}
}
