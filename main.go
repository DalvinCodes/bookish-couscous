package main

import (
	"github.com/DalvinCodes/bookish-couscous/config"
	"github.com/DalvinCodes/bookish-couscous/utils/errs"
	"github.com/DalvinCodes/bookish-couscous/utils/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	v, err := config.LoadConfig("config")
	errs.HandleFatalError(err)

	cnfg, err := config.ParseConfig(v)
	errs.HandleFatalError(err)

	logger.ConfigureLogger(cnfg)
	logger.Info("Logger Initialized")

	app := fiber.New()
	app.Use(logger.LoggingMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
