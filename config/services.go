package config

import "github.com/rizky201008/subscription-tracker-api/services"

var SettingService services.SettingService
var SubscriptionService services.SubscriptionService

func InitSettingService() {
	SettingService = services.NewSettingService(Db, SettingRepository)
}

func InitSubscriptionService() {
	SubscriptionService = services.NewSubscriptionService(Db, SubscriptionRepository)
}
