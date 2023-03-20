package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"fmt"
)

func registerUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := userStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	AddNewUser(payload.ID, payload.Email)
	return nil;
}

func lastID(c *fiber.Ctx) error {
	value := fmt.Sprintf("%s", GetLastID)
	valueToInt, err := strconv.Atoi(value);
	if err != nil {
		fmt.Println(err)
	}
	
	c.Write([]byte(strconv.Itoa(valueToInt)))
	return nil;
}


func main() {
	app := fiber.New()
	app.Post("/register", registerUser)
	app.Get("/lastid", lastID)
	app.Listen(":8080")
}