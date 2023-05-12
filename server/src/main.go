package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"fmt"
)

func registerUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := userStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	AddNewUser(payload.ID, payload.Email)
	return c.SendStatus(c.Response().StatusCode())
}

func lastID(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	fmt.Println("Katchau")
	c.Write([]byte(GetLastID()))
	return c.SendStatus(c.Response().StatusCode())}


func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))
	app.Post("/register", registerUser)
	app.Get("/lastid", lastID)
	app.Listen(":8080")
}