package book

import (
	"database/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func List(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func Retrieve(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func Create(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	c.BodyParser(book)
	db.Create(&book)
	return c.JSON(book)
}

func Update(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	updated_book := new(Book)
	c.BodyParser(updated_book)
	fmt.Print(updated_book)

	var book Book
	db.Find(&book, id)

	db.Model(&book).Updates(map[string]interface{}{"title": updated_book.Title, "Author": updated_book.Author, "rating": updated_book.Rating})
	return c.JSON(book)
}

func Destroy(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	return c.SendString("deleted")
}
