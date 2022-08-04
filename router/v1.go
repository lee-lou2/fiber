package router

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupV1Routes(api fiber.Router) {
	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Get("/verified/email", handler.AuthVerifiedEmail)
	auth.Post("/verified/email/re", handler.AuthReVerifiedEmail)
	auth.Post("/verified/phone", handler.AuthVerifiedPhone)
	auth.Post("/verified/phone/code", handler.AuthVerifiedPhoneCode)

	// User
	user := api.Group("/user")
	user.Post("/", handler.UserCreate)
	user.Post("/password/send", handler.UserSendChangePasswordEmail)
	user.Post("/password", handler.UserChangePassword)
}
