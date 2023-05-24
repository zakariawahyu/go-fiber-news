package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-news/modules/content/services"
	"github.com/zakariawahyu/go-fiber-news/utils/response"
)

type ContentController struct {
	contentServices services.ContentServices
}

func NewContentController(serv services.ContentServices) ContentController {
	return ContentController{
		contentServices: serv,
	}
}

func (ctrl *ContentController) Routes(route fiber.Router) {
	route.Get("/read/:slug", ctrl.GetBySlug)
}

func (ctrl *ContentController) GetBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	content := ctrl.contentServices.GetBySlug(slug)

	return ctx.JSON(response.NewSuccessResponse(fiber.StatusOK, content))
}
