package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewSampleRouter(router fiber.Router, db *gorm.DB) {
	ar := repositories.NewSampleRepository(db)
	as := services.NewSampleServices(ar)
	ac := controller.NewSampleController(as)
	sample_route := router.Group("/sample", middleware.AccessToken, func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	sample_route.Get("/all", middleware.AccessToken, ac.FindAll)
}
