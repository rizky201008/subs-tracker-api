package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/config"
)

func settingRoute(api fiber.Router) {
	setting := api.Group("/setting")
	controller := config.SettingController
	setting.Get("/", controller.GetSetting)
	setting.Post("/", controller.CreateSetting)
	setting.Put("/", controller.UpdateSetting)
}
