package main

import (
	"github.com/aldobarr/go-api-example/api/database"
	"github.com/aldobarr/go-api-example/api/handlers"
	"github.com/dgraph-io/badger/v4"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	database.DB = db

	app := fiber.New()

	handlers.InitValidators()
	setupRoutes(app)

	app.Listen(":8080")
}
