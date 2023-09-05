package controller

import "github.com/gofiber/fiber/v2"

func (ctrl *Controller) PingController(ctx *fiber.Ctx) error {
	res := ctrl.service.GetPing()
	return ctx.Status(fiber.StatusOK).JSON(res)
}
