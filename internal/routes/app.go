package routes

import (
	"time"

	"github.com/fnunezzz/go-project-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

type appRoute struct {
	app         fiber.Router
	appHealthService service.AppService
}

type AppRoute interface {
	Health()
}

type appDto struct {
	Time time.Time `json:"time"`
	Status bool `json:"status"`
}



func NewAppRoute(app fiber.Router, appHealthService service.AppService) {
	route := &appRoute{app: app, appHealthService: appHealthService}
	route.Health()
}

func (s *appRoute) Health() {
	s.app.Get("/status", func(c *fiber.Ctx) error {
		test, err := s.appHealthService.HealthCheck()
		dto := &appDto{Time: test, Status: true}
		status := 200
		if err != nil {
			dto.Status = false
			status = 500
		}
		return c.Status(status).JSON(dto)
	})

}