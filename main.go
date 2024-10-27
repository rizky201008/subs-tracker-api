package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/config"
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
}

func main() {
	app := fiber.New()

	router.MainRouter(app)

	err := app.Listen(config.Vipers.GetString("app.port"))
	if err != nil {
		panic(err)
	}
}
