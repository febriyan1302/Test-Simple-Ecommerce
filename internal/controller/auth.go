package controller

import (
	"github.com/febriyan1302/catalog-service/internal/entity"
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service service.Auth
}

type Auth interface {
	AuthController(ctx *fiber.Ctx) error
}

func ConstructAuthController(service service.Auth) *AuthController {
	return &AuthController{
		service: service,
	}
}

// Auth Route
func (ctrl *AuthController) AuthController(ctx *fiber.Ctx) error {
	authRequest := new(model.AuthRequest)
	if err := ctx.BodyParser(authRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	authResponse, err := ctrl.service.ValidateAuth(*authRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(entity.Response{
			Code:    ctx.Response().StatusCode(),
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(entity.Response{
		Code:    ctx.Response().StatusCode(),
		Message: entity.MessageSuccess,
		Data:    authResponse,
	})
}
