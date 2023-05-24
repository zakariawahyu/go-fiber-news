package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-news/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
