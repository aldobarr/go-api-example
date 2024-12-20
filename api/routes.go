package main

import (
	"github.com/aldobarr/go-api-example/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/receipts/:id/points", handlers.GetPoints)
	app.Post("/receipts/process", handlers.ProcessReceipt)
}
