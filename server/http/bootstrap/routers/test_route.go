package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type TestRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route TestRoutes) RegisterRoute() {
	handler := handlers.TestHandler{Handler: route.Handler}

	authenticationRoutes := route.RouteGroup.Group("/test")
	authenticationRoutes.Get("", handler.TestProcess)
}

