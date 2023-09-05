package controller

import (
	"github.com/febriyan1302/catalog-service/internal/entity"
	"github.com/febriyan1302/catalog-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service service.Product
}

type Product interface {
	ListProduct(ctx *fiber.Ctx) error
	DetailProduct(ctx *fiber.Ctx) error
}

func ConstructorProductController(service service.Product) *ProductController {
	return &ProductController{service: service}
}

func (ctrl *ProductController) ListProduct(ctx *fiber.Ctx) error {
	product, err := ctrl.service.FindAllProduct()
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
		Data:    product,
	})
}

func (ctrl *ProductController) DetailProduct(ctx *fiber.Ctx) error {
	product, err := ctrl.service.GetProductDetail(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(entity.Response{
			Code:    ctx.Response().StatusCode(),
			Message: entity.MessageFailed,
			Data:    nil,
		})
	}

	if product.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(entity.Response{
			Code:    ctx.Response().StatusCode(),
			Message: entity.MessageNotFound,
			Data:    nil,
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(entity.Response{
		Code:    ctx.Response().StatusCode(),
		Message: entity.MessageSuccess,
		Data:    product,
	})

}
