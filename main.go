package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dhaliwal-h/go-fiber-crm-basic/database"
	"github.com/dhaliwal-h/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
	app.Get("/", func(c *fiber.Ctx) {
		c.Fasthttp.Response.Header.Set("Content-Type", "application/json")

		json.NewEncoder(c.Fasthttp.Response.BodyWriter()).Encode("Hello World.")
	})
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "lead.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main() {
	fmt.Println("Hello world")
	initDatabase()
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
