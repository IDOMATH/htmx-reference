package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
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

func main() {

	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("GET /", handleHome)

	fmt.Println("server started on port 8080")
	log.Fatal(server.ListenAndServe())
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home"))
}
