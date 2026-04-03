package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	name := "John"

	fmt.Println("Hello, " + name)
	app := fiber.New()
	fmt.Print(app)
}
