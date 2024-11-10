package exception

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/models/web"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	notFoundErr, notFound := notFoundError(ctx, err)
	if notFound {
		return notFoundErr
	}
	unauthorizedErr, unauthorized := unauthorizedError(ctx, err)
	if unauthorized {
		return unauthorizedErr
	}

	ctx.Status(fiber.StatusInternalServerError)
	return ctx.JSON(
		web.Response{
			StatusCode: web.RcError,
			Message:    err.Error(),
			Data:       nil,
		},
	)
}

func notFoundError(ctx *fiber.Ctx, err error) (error, bool) {
	var exception NotFoundError
	ok := errors.As(err, &exception)
	if ok {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(web.Response{
			StatusCode: web.RcNotFound,
			Message:    exception.Error(),
			Data:       nil,
		}), true

	} else {
		return nil, false
	}
}

func unauthorizedError(ctx *fiber.Ctx, err error) (error, bool) {
	var exception UnauthorizedError
	ok := errors.As(err, &exception)
	if ok {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(web.Response{
			StatusCode: web.RcUnauthorized,
			Message:    exception.Error(),
			Data:       nil,
		}), true

	} else {
		return nil, false
	}
}
