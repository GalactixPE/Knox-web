package main

import (
	"context"
	"github.com/Mohagames205/Knox/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		AppName:                 "Knox",
		Views:                   engine,
		EnableTrustedProxyCheck: true,
	})

	_ = godotenv.Load()

	app.Static("/cdn", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("verify", fiber.Map{})
	})

	app.Post("/api/whitelist", func(ctx *fiber.Ctx) error {
		payload := struct {
			Name string `json:"username"`
		}{}

		if err := ctx.BodyParser(&payload); err != nil {
			return err
		}

		db.Client().Set(context.Background(), payload.Name, ctx.IP(), 0)

		return ctx.JSON(fiber.Map{"message": "Ok, IP queued for whitelisting"})
	})

	log.Fatal(app.Listen("0.0.0.0:3000"))

}
