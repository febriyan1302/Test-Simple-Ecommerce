package controller

import (
	"github.com/febriyan1302/catalog-service/internal/entity"
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/service"
	"github.com/febriyan1302/catalog-service/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type CartController struct {
	service service.Cart
}

type Cart interface {
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
}

func ConstructorCartController(service service.Cart) *CartController {
	return &CartController{service: service}
}

func (ctrl *CartController) Create(ctx *fiber.Ctx) error {
	deserializeToken := utils.DeserializeToken(ctx)

	cartRequest := new(model.CartRequest)
	cartRequest.UserId = deserializeToken.ID
	if err := ctx.BodyParser(cartRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := ctrl.service.CreateCart(*cartRequest)
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
		Data:    nil,
	})
}

func (ctrl *CartController) Get(ctx *fiber.Ctx) error {
	deserializeToken := utils.DeserializeToken(ctx)

	cartData, err := ctrl.service.GetCart(deserializeToken.ID)
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
		Data:    cartData,
	})
}
