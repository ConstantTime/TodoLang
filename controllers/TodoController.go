package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"todo/db"
	"todo/models"
)

var dbClient *sqlx.DB

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Adding todo: %v\n", todo)

	_ = db.AddTodo(dbClient, todo)

	res, err := json.Marshal(todo)
	if err != nil {
		log.Printf("Error while marshalling AddTodo response: %v", err)
	}

	returnJsonResponse(w, res)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	todo, _ := db.GetTodoById(dbClient, id)

	res, err := json.Marshal(todo)
	if err != nil {
		log.Printf("Error while marshalling GetTodo response: %v", err)
	}

	returnJsonResponse(w, res)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, _ := db.GetTodos(dbClient)

	res, err := json.Marshal(todos)
	if err != nil {
		log.Printf("Error while marshalling GetTodos response: %v", err)
	}

	returnJsonResponse(w, res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := db.DeleteTodoById(dbClient, id)

	if err != nil {
		log.Printf("Error while deleting todo (id = %v): %v", id, err)
	}
}

func returnJsonResponse(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func HandleAllTodoRequests() {
	dbClient = db.ConnectToDatabase()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos", GetTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8081", router))
}
