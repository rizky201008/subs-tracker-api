package config

import "github.com/rizky201008/subscription-tracker-api/middleware"

var Middleware middleware.Middleware

func InitMiddleware() {
	Middleware = middleware.NewMiddleware(Vipers, Db, Firebase)
}
