package bootstrap

import (
	"github.com/adamsh231/majoo/server/http/bootstrap/routers"
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func (boot Bootstrap) RegisterRoute() {
	handlerType := handlers.Handler{
		UcContract: &boot.UcContract,
	}

	// HealthCheck Route
	rootParentGroup := boot.UcContract.App.Group("/api")
	rootParentGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("Work!")
	})

	// Route Group
	apiV1 :=rootParentGroup.Group("/v1")

	// Auth route
	AuthRoute := routers.AuthRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	AuthRoute.RegisterRoute()

}
