package controller

import (
	"github.com/febriyan1302/catalog-service/internal/entity"
	"github.com/febriyan1302/catalog-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	service service.Category
}

type Category interface {
	ListCategory(ctx *fiber.Ctx) error
}

func ConstructorCategoryController(service service.Category) *CategoryController {
	return &CategoryController{service: service}
}

func (ctrl *CategoryController) ListCategory(ctx *fiber.Ctx) error {
	categories, err := ctrl.service.FindAllCategory()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(entity.Response{
			Code:    ctx.Response().StatusCode(),
			Message: entity.MessageFailed,
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(entity.Response{
		Code:    ctx.Response().StatusCode(),
		Message: entity.MessageSuccess,
		Data:    categories,
	})

}
