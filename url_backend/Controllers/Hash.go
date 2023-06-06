package Controller

import (
    /*
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    */
	"github.com/gofiber/fiber/v2"
)

//Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

func GetHashedUrl(c *fiber.Ctx) error {
    id := c.Params("id")
    return c.SendString("Create hashed URL " + id)
}

func CreateHashedUrlV1(c *fiber.Ctx) error {
    return c.SendString("Create hashed URL")
}

//AddBook
/*
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&book)

	return c.Status(200).JSON(book)
}
*/
