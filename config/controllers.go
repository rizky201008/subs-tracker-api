package config

import "github.com/rizky201008/subscription-tracker-api/controllers"

var SettingController controllers.SettingController
var SubscriptionController controllers.SubscriptionController

func InitSettingController() {
	SettingController = controllers.NewSettingController(SettingService)
	SubscriptionController = controllers.NewSubscriptionController(SubscriptionService)
}
