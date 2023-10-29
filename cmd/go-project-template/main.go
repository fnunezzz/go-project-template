package main

import (
	"github.com/fnunezzz/go-project-template/internal/repository"
	"github.com/fnunezzz/go-project-template/internal/routes"
	"github.com/fnunezzz/go-project-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
	api := app.Group("/api")
	dao := repository.InitConn()

	// Services
	appService := service.NewAppService(dao)

	// Routes
	routes.NewAppRoute(api, appService)
	
    app.Listen(":4000")
}