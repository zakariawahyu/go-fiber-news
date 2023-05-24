package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-news/helpers"
	"github.com/zakariawahyu/go-fiber-news/utils/response"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	code := helpers.GetStatusCode(err)
	return ctx.Status(code).JSON(response.ErrorResponse{
		Success: false,
		Code:    code,
		Errors:  err.Error(),
	})
}
