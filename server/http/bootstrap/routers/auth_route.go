package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route AuthRoutes) RegisterRoute() {
	handler := handlers.AuthHandler{Handler: route.Handler}

	authenticationRoutes := route.RouteGroup.Group("/login")
	authenticationRoutes.Get("", handler.Login)
}

