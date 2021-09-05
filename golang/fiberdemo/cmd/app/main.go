package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		defer func() {
			c.Response().Header.Add("x-Response-Time", strconv.Itoa(int(time.Since(start).Microseconds())))
		}()
		return c.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(200 * time.Microsecond)
		return c.SendString("Hello, World 👋!")
	})
	app.Listen("localhost:8999")
}
