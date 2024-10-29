package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/helper"
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"github.com/rizky201008/subscription-tracker-api/models/web"
	"github.com/rizky201008/subscription-tracker-api/repositories"
	"gorm.io/gorm"
)

type SubscriptionService interface {
	GetAll(ctx *fiber.Ctx) []web.SubscriptionResponse
	GetById(ctx *fiber.Ctx) web.SubscriptionResponse
	Create(ctx *fiber.Ctx) web.SubscriptionResponse
	Update(ctx *fiber.Ctx) web.SubscriptionResponse
}

type SubscriptionServiceImpl struct {
	Db               *gorm.DB
	SubscriptionRepo repositories.SubscriptionRepository
}

func (service SubscriptionServiceImpl) GetAll(ctx *fiber.Ctx) []web.SubscriptionResponse {
	var response []web.SubscriptionResponse
	result, err := service.SubscriptionRepo.GetAll(service.Db)
	if err != nil {
		panic(err)
	}
	response = helper.SubscriptionToResponses(result)
	return response
}

func (service SubscriptionServiceImpl) GetById(ctx *fiber.Ctx) web.SubscriptionResponse {
	id := ctx.Params("id")
	var response web.SubscriptionResponse
	result, err := service.SubscriptionRepo.GetByID(id, service.Db)
	if err != nil {
		panic(err)
	}
	response = helper.SubscriptionToResponse(result)
	return response
}

func (service SubscriptionServiceImpl) Create(ctx *fiber.Ctx) web.SubscriptionResponse {
	p := new(web.SubscriptionRequest)
	if err := ctx.BodyParser(p); err != nil {
		panic(err)
	}
	var response web.SubscriptionResponse

	userId := ctx.GetRespHeader("userId")
	data := domain.Subscription{
		Amount:       p.Amount,
		DueDate:      p.DueDate,
		UserId:       userId,
		ColorHex:     p.ColorHex,
		Cycle:        domain.StringToBilling(p.Cycle),
		PlatformName: p.PlatformName,
		Reminder:     p.Reminder,
	}
	result, err := service.SubscriptionRepo.Save(data, service.Db)
	if err != nil {
		panic(err)
	}
	response = helper.SubscriptionToResponse(result)
	return response
}

func (service SubscriptionServiceImpl) Update(ctx *fiber.Ctx) web.SubscriptionResponse {
	id := ctx.Params("id")
	p := new(web.SubscriptionRequest)
	var response web.SubscriptionResponse
	if err := ctx.BodyParser(p); err != nil {
		panic(err)
	}
	result, err := service.SubscriptionRepo.GetByID(id, service.Db)
	if err != nil {
		panic(err)
	}
	result.Reminder = p.Reminder
	result.Cycle = domain.StringToBilling(p.Cycle)
	result.ColorHex = p.ColorHex
	result.PlatformName = p.PlatformName
	result.DueDate = p.DueDate
	result.Amount = p.Amount

	updated, err := service.SubscriptionRepo.Update(result, service.Db)
	if err != nil {
		panic(err)
	}
	response = helper.SubscriptionToResponse(updated)
	return response
}

func NewSubscriptionService(db *gorm.DB, subRepo repositories.SubscriptionRepository) SubscriptionService {
	return SubscriptionServiceImpl{
		Db:               db,
		SubscriptionRepo: subRepo,
	}
}
