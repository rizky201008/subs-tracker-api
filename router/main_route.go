package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/models/web"
)

func MainRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})
	api := app.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(web.Response{
			StatusCode: web.RcSuccess,
			Data:       nil,
			Message:    "Sukses",
		})
	})
	settingRoute(api)
	subscriptionRoute(api)
}
