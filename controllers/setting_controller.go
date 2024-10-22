package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/models/web"
	"github.com/rizky201008/subscription-tracker-api/services"
)

type SettingController interface {
	GetSetting(ctx *fiber.Ctx) error
	UpdateSetting(ctx *fiber.Ctx) error
	CreateSetting(ctx *fiber.Ctx) error
}

type SettingControllerImpl struct {
	SettingService services.SettingService
}

func NewSettingController(settingService services.SettingService) SettingControllerImpl {
	return SettingControllerImpl{
		SettingService: settingService,
	}
}

func (controller SettingControllerImpl) GetSetting(ctx *fiber.Ctx) error {
	response := web.NewResponse()
	response.Data = controller.SettingService.GetSetting(ctx)

	return ctx.JSON(response)
}

func (controller SettingControllerImpl) UpdateSetting(ctx *fiber.Ctx) error {
	response := web.NewResponse()
	response.Data = controller.SettingService.UpdateSetting(ctx)
	return ctx.JSON(response)
}

func (controller SettingControllerImpl) CreateSetting(ctx *fiber.Ctx) error {
	response := web.NewResponse()
	response.Data = controller.SettingService.CreateSetting(ctx)
	return ctx.JSON(response)
}
