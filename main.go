package main

import (
	"todo/controllers"
	"todo/db/migrate"
)

func main() {
	migrate.RunMigrations()
	controllers.HandleAllTodoRequests()
}
