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

	data = &types.TemplateData{Count: 0}

	router.HandleFunc("GET /", handleHome)
	router.HandleFunc("POST /count", handlePostCount)

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
