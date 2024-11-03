package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/config"
)

func subscriptionRoute(api fiber.Router) {
	setting := api.Group("/subscriptions")
	controller := config.SubscriptionController
	setting.Get("/", controller.GetSubscriptions)
	setting.Post("/", controller.CreateSubscription)
	setting.Put("/", controller.UpdateSubscription)
}
