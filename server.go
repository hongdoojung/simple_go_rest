package main

import (
	"book/book"
	"database/database"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Seoul",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Print("Database connected")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Print("Database migrated")
}

func setUpRoutes(app *fiber.App) {
	app.Get("api/v1/book", book.List)
	app.Post("api/v1/book", book.Create)
	app.Get("api/v1/book/:id", book.Retrieve)
	app.Put("api/v1/book/:id", book.Update)
	app.Delete("api/v1/book/:id", book.Destroy)
	return
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Get("/", helloWorld)
	app.Listen(":3000")
}
