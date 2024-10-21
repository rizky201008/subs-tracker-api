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
