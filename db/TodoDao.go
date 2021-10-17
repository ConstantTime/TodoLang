package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"todo/models"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sqlx.DB {
	connStr := "user=postgres dbname=pg_migrations_example sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AddTodo(dbClient *sqlx.DB, todo models.Todo) error {
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

func GetTodoById(dbClient *sqlx.DB, id string) (models.Todo, error) {
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

func GetTodos(dbClient *sqlx.DB) ([]models.Todo, error) {
	var todoEntities []models.TodoEntity
	err := dbClient.Select(&todoEntities,`select id, title, description from todos where isdeleted=false`)

	if err != nil {
		log.Printf("Error while getting all todos from DB: %v", err)
	}

	var todos []models.Todo
	for i := range todoEntities {
		todos = append(todos, models.Todo{
			Id:          todoEntities[i].Id,
			Title:       todoEntities[i].Title,
			Description: todoEntities[i].Description,
		})
	}

	return todos, err
}

func DeleteTodoById(dbClient *sqlx.DB, id string) error {
	_, err := dbClient.Query(`update todos set isDeleted=true where id=$1`, id)

	if err != nil {
		log.Printf("Error while get todo task with id=%v %v", id, err)
	}

	return err
}
