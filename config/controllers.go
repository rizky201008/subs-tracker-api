package config

import "github.com/rizky201008/subscription-tracker-api/controllers"

var SettingController controllers.SettingController

func InitSettingController() {
	SettingController = controllers.NewSettingController(SettingService)
}
