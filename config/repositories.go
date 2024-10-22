package config

import "github.com/rizky201008/subscription-tracker-api/repositories"

var SettingRepository repositories.SettingRepository

func InitSettingRepository() {
	SettingRepository = repositories.NewSettingRepository()
}
