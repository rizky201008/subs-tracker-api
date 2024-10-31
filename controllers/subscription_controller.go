package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/models/web"
	"github.com/rizky201008/subscription-tracker-api/services"
)

type SubscriptionController interface {
	GetSubscriptions(ctx *fiber.Ctx) error
	GetSubscription(ctx *fiber.Ctx) error
	CreateSubscription(ctx *fiber.Ctx) error
	UpdateSubscription(ctx *fiber.Ctx) error
}

type SubscriptionControllerImpl struct {
	SubscriptionService services.SubscriptionService
}

func (controller SubscriptionControllerImpl) GetSubscriptions(ctx *fiber.Ctx) error {
	var response = web.NewResponse()
	response.Data = controller.GetSubscriptions(ctx)
	return ctx.JSON(response)
}

func (controller SubscriptionControllerImpl) GetSubscription(ctx *fiber.Ctx) error {
	var response = web.NewResponse()
	response.Data = controller.GetSubscription(ctx)
	return ctx.JSON(response)
}

func (controller SubscriptionControllerImpl) CreateSubscription(ctx *fiber.Ctx) error {
	var response = web.NewResponse()
	response.Data = controller.CreateSubscription(ctx)
	return ctx.JSON(response)
}

func (controller SubscriptionControllerImpl) UpdateSubscription(ctx *fiber.Ctx) error {
	var response = web.NewResponse()
	response.Data = controller.UpdateSubscription(ctx)
	return ctx.JSON(response)
}

func NewSubscriptionController(subscriptionService services.SubscriptionService) SubscriptionController {
	return SubscriptionControllerImpl{
		SubscriptionService: subscriptionService,
	}
}
