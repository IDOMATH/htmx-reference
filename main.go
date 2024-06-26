package main

import (
	"fmt"
	"log"
	"net/http"
)

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
