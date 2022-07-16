package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func setUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()
	setUpRoutes(app)
}
