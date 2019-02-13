package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// A minimum in-memory mock storage, it's not thread safe!
var tasks []Task

func main() {
	const Port = ":3000"

	routes := mux.NewRouter()

	routes.Methods("GET").Path("/tasks").HandlerFunc(GetTasks)
	routes.Methods("GET").Path("/tasks/{id}").HandlerFunc(GetTask)
	routes.Methods("POST").Path("/tasks").HandlerFunc(CreateTask)

	log.Printf("Server running on %s", Port)
	log.Fatal(http.ListenAndServe(Port, routes))
}

// GET /tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GET /tasks/{id}
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	for _, item := range tasks {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) // no matches found
}

// POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
