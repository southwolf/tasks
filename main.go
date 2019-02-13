package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const Port = ":3000"

	routes := mux.NewRouter()

	routes.HandleFunc("/tasks", TasksHandler)

	log.Printf("Server running on %s", Port)
	log.Fatal(http.ListenAndServe(Port, routes))
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You've reached tasks routes")
}
