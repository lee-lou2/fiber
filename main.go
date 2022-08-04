package main

import (
	"log"

	"api/database"
	"api/middleware"
	"api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "CAR App v1.0.1",
	})

	// 데이터베이스 연결
	database.ConnectDB()

	// 미들 웨어
	middleware.SetupMiddleWare(app)

	// 라우터
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":80"))
}
