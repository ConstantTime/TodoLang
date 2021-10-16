package main

import (
	"todo/db/migrate"
)

func main() {
	migrate.RunMigrations()
}
