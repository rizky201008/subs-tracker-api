package helper

import (
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"github.com/rizky201008/subscription-tracker-api/models/web"
)

func SettingToResponse(data domain.Setting) web.SettingResponse {
	return web.SettingResponse{
		UserID:   data.UserID,
		Currency: data.Currency,
	}
}

func SubscriptionToResponses(data []domain.Subscription) []web.SubscriptionResponse {
	var result []web.SubscriptionResponse
	for _, value := range data {
		result = append(result, web.SubscriptionResponse{
			PlatformName: value.PlatformName,
			Amount:       value.Amount,
			DueDate:      value.DueDate,
			Reminder:     value.Reminder,
			Cycle:        value.Cycle.Value(),
			ColorHex:     value.ColorHex,
		})
	}
	return result
}

func SubscriptionToResponse(data domain.Subscription) web.SubscriptionResponse {
	return web.SubscriptionResponse{
		PlatformName: data.PlatformName,
		Amount:       data.Amount,
		DueDate:      data.DueDate,
		Reminder:     data.Reminder,
		Cycle:        data.Cycle.Value(),
		ColorHex:     data.ColorHex,
	}
}
