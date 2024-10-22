package config

import "github.com/rizky201008/subscription-tracker-api/services"

var SettingService services.SettingService

func InitSettingService() {
	SettingService = services.NewSettingService(Db, SettingRepository)
}
