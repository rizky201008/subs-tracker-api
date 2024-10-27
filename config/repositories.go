package config

import "github.com/rizky201008/subscription-tracker-api/repositories"

var SettingRepository repositories.SettingRepository
var SubscriptionRepository repositories.SubscriptionRepository

func InitSettingRepository() {
	SettingRepository = repositories.NewSettingRepository()
}

func InitSubscriptionRepository() {
	SubscriptionRepository = repositories.NewSubscriptionRepository()
}
