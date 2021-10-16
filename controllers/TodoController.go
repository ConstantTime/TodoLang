package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo/db"
	"todo/models"
)

var dbClient *sql.DB

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Adding todo: %v\n", todo)

	_ = db.AddTodo(dbClient, todo)
	json.NewEncoder(w).Encode(todo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	todo, _ := db.GetTodoById(dbClient, id)
	json.NewEncoder(w).Encode(todo)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, _ := db.GetTodos(dbClient)
	json.NewEncoder(w).Encode(todos)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := db.DeleteTodoById(dbClient, id)

	if err != nil {
		log.Printf("Error while deleting todo (id = %v): %v", id, err)
	}
}

func HandleAllTodoRequests() {
	dbClient = db.ConnectToDatabase()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", addTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", getTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos", getTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", deleteTodo).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8081", router))
}
