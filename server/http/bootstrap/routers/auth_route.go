package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/adamsh231/majoo/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route AuthRoutes) RegisterRoute() {
	handler := handlers.AuthHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract:route.Handler.UcContract}

	authenticationRoutes := route.RouteGroup.Group("/login")
	authenticationRoutes.Use(jwtMiddleware.New)

	authenticationRoutes.Get("", handler.Login)
}

