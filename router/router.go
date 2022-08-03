package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func healthCheckHandler(c *fiber.Ctx) error {
	// 헬스 체크
	return c.JSON(fiber.Map{"isRunning": true})
}

func SetupRoutes(app *fiber.App) {
	// Monitoring
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	app.Get("/health", healthCheckHandler)
	// 정적 파일
	app.Static("/static", "./static")
	app.Static("/views", "./views")

	// Version
	api := app.Group("/v1", logger.New())
	SetupV1Routes(api)
}
