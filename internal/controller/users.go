package controller

import (
	"github.com/febriyan1302/catalog-service/internal/entity"
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service service.User
}

type User interface {
	CreateController(ctx *fiber.Ctx) error
}

func ConstructUserController(service service.User) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctrl *UserController) CreateController(ctx *fiber.Ctx) error {
	userRequest := new(model.UserRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userResponse, err := ctrl.service.CreateService(*userRequest)
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
		Data:    userResponse,
	})
}
