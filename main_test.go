package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestTasksHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(TasksHandler)
	handler.ServeHTTP(response, request)

	expected := "You've reached tasks routes"
	actual := response.Body.String()
	if actual != expected {
		t.Errorf("Wrong response: got %v want %v", actual, expected)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
