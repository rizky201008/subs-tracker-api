package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/helper"
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"github.com/rizky201008/subscription-tracker-api/models/web"
	"github.com/rizky201008/subscription-tracker-api/repositories"
	"gorm.io/gorm"
)

type SettingService interface {
	GetSetting(ctx *fiber.Ctx) web.SettingResponse
	UpdateSetting(ctx *fiber.Ctx) web.SettingResponse
	CreateSetting(ctx *fiber.Ctx) web.SettingResponse
}

type SettingServiceImpl struct {
	Db          *gorm.DB
	SettingRepo repositories.SettingRepository
}

func NewSettingService(db *gorm.DB, settingRepo repositories.SettingRepository) SettingService {
	return SettingServiceImpl{
		Db:          db,
		SettingRepo: settingRepo,
	}
}

func (service SettingServiceImpl) GetSetting(ctx *fiber.Ctx) web.SettingResponse {
	userId := ctx.GetRespHeader("userId")
	var response web.SettingResponse

	setting, err := service.SettingRepo.GetSetting(userId, service.Db)
	if err != nil {
		panic(err)
	}
	response = helper.SettingToResponse(setting)
	return response
}

func (service SettingServiceImpl) UpdateSetting(ctx *fiber.Ctx) web.SettingResponse {
	p := new(web.SettingRequest)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	userId := ctx.GetRespHeader("userId")
	var response web.SettingResponse

	setting, err := service.SettingRepo.GetSetting(userId, service.Db)
	if err != nil {
		panic(err)
	}
	setting.Currency = p.Currency
	settingUpdated, err := service.SettingRepo.UpdateSetting(setting, service.Db)
	response = helper.SettingToResponse(settingUpdated)
	return response
}

func (service SettingServiceImpl) CreateSetting(ctx *fiber.Ctx) web.SettingResponse {
	p := new(web.SettingRequest)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	userId := ctx.GetRespHeader("userId")
	var response web.SettingResponse
	setting := domain.Setting{
		UserID:   userId,
		Currency: p.Currency,
	}

	getSetting, err := service.SettingRepo.GetSetting(userId, service.Db)
	if err == nil {
		response = helper.SettingToResponse(getSetting)
	}

	saveSetting, err := service.SettingRepo.SaveSetting(setting, service.Db)
	if err != nil {
		panic(err)
	}
	response = helper.SettingToResponse(saveSetting)
	return response
}
