package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/idomath/htmx-reference/middleware"
	"github.com/idomath/htmx-reference/render"
	"github.com/idomath/htmx-reference/types"
)

var data *types.TemplateData

func main() {

	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(router),
	}

	data = &types.TemplateData{Count: 0,
		Contacts: []types.Contact{
			{Name: "John", Email: "jd@gmail.com"},
			{Name: "Clara", Email: "cd@gmail.com"},
		}}

	data.Contacts = []types.Contact{
		{Name: "John", Email: "jd@gmail.com"},
		{Name: "Clara", Email: "cd@gmail.com"},
	}

	router.HandleFunc("GET /", handleHome)
	router.HandleFunc("POST /count", handlePostCount)
	router.HandleFunc("POST /contacts", handlePostContacts)

	fmt.Println("server started on port 8080")
	log.Fatal(server.ListenAndServe())
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "index.html", data)
}

func handlePostCount(w http.ResponseWriter, r *http.Request) {
	data.Count++
	render.Template(w, r, "count.html", data)
}

func handlePostContacts(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	data.Contacts = append(data.Contacts, types.Contact{Name: name, Email: email})
	render.Template(w, r, "display.html", data)
}
