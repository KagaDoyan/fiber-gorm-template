package route

import (
	"fmt"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, fb *firebase.App, db *gorm.DB) {
	api := app.Group("/api/v1", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})

	api.Post("/health", func(c *fiber.Ctx) error {
		hostname, err := os.Hostname()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Error getting hostname: %s", err),
			})
		}
		currentTime := time.Now().Format(time.RFC3339)
		if err := c.Status(200).JSON(fiber.Map{
			"hostname":  hostname,
			"timestamp": currentTime,
			"msg":       "Connect OK...!",
		}); err != nil {
			return err
		}
		return nil
	})

	NewSampleRouter(api, db)
}
