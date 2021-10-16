package db

import (
	"database/sql"
	"fmt"
	"log"
	"todo/models"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	connStr := "user=postgres dbname=pg_migrations_example sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AddTodo(dbClient *sql.DB, todo models.Todo) error {
	_, err := dbClient.Query(`insert into 
    	todos(id, title, description, isDeleted) values($1, $2, $3, false)`,
		todo.Id,
		todo.Title,
		todo.Description,
	)

	if err != nil {
		log.Printf("Error while adding todo task to db: %v", err)
	}

	return err
}

func GetTodoById(dbClient *sql.DB, id string) (models.Todo, error) {
	var todo models.TodoEntity
	err := dbClient.QueryRow(`select id, title, description from todos where id=$1`, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
	)

	if err != nil {
		log.Printf("Error while get todo with id=%v %v", id, err)
	}

	return models.Todo{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
	}, err
}

func GetTodos(dbClient *sql.DB) ([]models.Todo, error) {
	rows, err := dbClient.Query(`select id, title from todos where isDeleted=false`)

	if err != nil {
		log.Printf("Error while getting all todos %v", err)
	}

	var todos []models.Todo

	for rows.Next() {
		fmt.Printf("rows are %v\n", rows)
		var todoEntity models.TodoEntity
		err = rows.Scan(&todoEntity)
		fmt.Printf("entity is %v\n", todoEntity)
		todos = append(todos, models.Todo{
			Id:          todoEntity.Id,
		})
	}

	return todos, err
}

func DeleteTodoById(dbClient *sql.DB, id string) error {
	_, err := dbClient.Query(`update todos set isDeleted=true where id=$1`, id)

	if err != nil {
		log.Printf("Error while get todo task with id=%v %v", id, err)
	}

	return err
}
