package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-fiber-news/config"
	"github.com/zakariawahyu/go-fiber-news/modules/content/controller"
	"log"
)

func NewAppHandler(app *fiber.App) {
	cfg := config.NewConfig()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"app":         cfg.App.Name,
			"version":     cfg.App.Version,
			"app_timeout": cfg.App.ContextTimeout,
		})
	})
}
func NewHandler(cfg *config.Config, serv *Services) {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	v1 := app.Group("/v1")

	NewAppHandler(app)
	contentCtrl := controller.NewContentController(serv.contentServices)
	contentCtrl.Routes(v1)

	log.Fatal(app.Listen(viper.GetString("APP_ADDRESS")))
}
