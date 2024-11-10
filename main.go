package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizky201008/subscription-tracker-api/config"
	"github.com/rizky201008/subscription-tracker-api/exception"
	"github.com/rizky201008/subscription-tracker-api/router"
)

func init() {
	config.InitViper()
	config.InitDb()
	config.InitSettingRepository()
	config.InitSubscriptionRepository()
	config.InitSettingService()
	config.InitSubscriptionService()
	config.InitSettingController()
	config.InitFirebase()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(recover.New())

	router.MainRouter(app)

	err := app.Listen(config.Vipers.GetString("app.port"))
	if err != nil {
		panic(err)
	}
}
