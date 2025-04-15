// internal/handlers/url.go
package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	app.Post("/shorten", shortenURL(db))
	app.Get("/:short", resolveURL(db))
}

func shortenURL(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Add URL shortening logic
		return c.SendString("Shorten URL endpoint")
	}
}

func resolveURL(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Add URL resolution logic
		return c.SendString("Resolve URL endpoint")
	}
}
