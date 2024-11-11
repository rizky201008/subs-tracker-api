package middleware

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/subscription-tracker-api/exception"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Middleware interface {
	RequireAuth(c *fiber.Ctx) error
}

type Impl struct {
	Viper    *viper.Viper
	DB       *gorm.DB
	Firebase *firebase.App
}

func (middleware Impl) RequireAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("X-Access-Token")
	if tokenString == "" {
		panic(exception.UnauthorizedError{
			Err: "Authorization error, token not found!",
		})
	}
	conteks := context.Background()
	client, err := middleware.Firebase.Auth(conteks)
	if err != nil {
		panic(err)
	}
	token, err := client.VerifyIDToken(conteks, tokenString)
	if err != nil {
		panic(exception.UnauthorizedError{Err: err.Error()})
	}
	ctx.Set("uid", token.UID)
	return ctx.Next()
}

func NewMiddleware(viper *viper.Viper, db *gorm.DB, fireBase *firebase.App) Middleware {
	return Impl{Viper: viper, DB: db, Firebase: fireBase}
}
