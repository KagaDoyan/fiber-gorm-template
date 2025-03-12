package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
)

type sampleController struct {
	sampleService services.SampleService
}

type SampleController interface {
	FindAll(ctx *fiber.Ctx) error
}

func (sampleController *sampleController) FindAll(ctx *fiber.Ctx) error {
	result, err := sampleController.sampleService.FindAll()
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func NewSampleController(sampleService services.SampleService) SampleController {
	return &sampleController{sampleService: sampleService}
}
