package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

func SetupMiddleWare(app *fiber.App) {
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// CSRF
	//app.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "header:X-Csrf-Token",
	//	CookieName:     "csrf_",
	//	CookieSameSite: "Strict",
	//	Expiration:     1 * time.Hour,
	//	KeyGenerator:   utils.UUID,
	//}))
	// 파비콘
	app.Use(favicon.New(favicon.Config{
		File: "./static/public/favicon.ico",
	}))
	// 리미터
	app.Use(limiter.New(limiter.Config{
		Max:        1,
		Expiration: 1 * time.Second,
		Next: func(c *fiber.Ctx) bool {
			// 우선은 전체 제외
			return true
		},
	}))
	// 로거
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Seoul",
	}))
	// PPROF
	app.Use(pprof.New())
	// panic 을 api 로 반환
	app.Use(recover.New())
}
