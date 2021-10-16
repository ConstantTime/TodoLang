package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/models"
)

type Todos []models.Todo

func allTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All todos endpoint hit\n")
	todos := Todos{
		models.Todo{
			Id:          "1",
			Title:       "test",
			Description: "test",
		},
	}

	json.NewEncoder(w).Encode(todos)
}

func HandleAllTodoRequests() {
	http.HandleFunc("/todos", allTodos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
