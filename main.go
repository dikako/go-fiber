package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Prefork:      true,
	})

	// prefix /api will be use middleware, when not define prefix middleware will be use for all request
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Middleware Before request")
		err := c.Next()
		fmt.Println("Middleware After request")
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	app.Get("/api/v1", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	if fiber.IsChild() {
		fmt.Println("I'm Child process")
	} else {
		fmt.Println("I'm Parent process")
	}
	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
