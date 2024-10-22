package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/config"
	"github.com/rizky201008/subscription-tracker-api/controllers"
)

func settingRoute(api fiber.Router) {
	setting := api.Group("/setting")
	controller := controllers.NewSettingController(config.SettingService)
	setting.Get("/", controller.GetSetting)
	setting.Post("/", controller.CreateSetting)
	setting.Put("/", controller.UpdateSetting)
}
