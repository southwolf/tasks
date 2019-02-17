package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTasks(t *testing.T) {
	taskPayload := `{"id":1,"title":"Hello world!","done":false}`
	request, _ := http.NewRequest("POST", "/tasks", strings.NewReader(taskPayload))
	response := httptest.NewRecorder()
	CreateTask(response, request)

	expected := taskPayload
	actual := strings.TrimSpace(response.Body.String())
	if actual != expected {
		t.Errorf("Wrong response: got %v want %v", actual, expected)
	}
}

func TestGetTasks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tasks", nil)
	response := httptest.NewRecorder()
	GetTasks(response, request)

	expected := `[{"id":1,"title":"Hello world!","done":false}]`
	actual := strings.TrimSpace(response.Body.String())
	if actual != expected {
		t.Errorf("Wrong response: got %v want %v", actual, expected)
	}
}

func TestGetTask(t *testing.T) {
	routes := mux.NewRouter()
	routes.HandleFunc("/tasks/{id}", GetTask)

	taskPayload := `{"id":1,"title":"Hello world!","done":false}`
	request, _ := http.NewRequest("GET", "/tasks/1", nil)
	response := httptest.NewRecorder()
	routes.ServeHTTP(response, request)

	expected := taskPayload
	actual := strings.TrimSpace(response.Body.String())
	if actual != expected {
		t.Errorf("Wrong response: got %v want %v", actual, expected)
	}
}
