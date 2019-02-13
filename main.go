package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const Port = ":3000"
	http.HandleFunc("/", greet)
	http.HandleFunc("/tasks", tasksHandler)
	log.Printf("Server running on %s", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You've reached tasks routes")
}
