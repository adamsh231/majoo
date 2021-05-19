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
	authHandler := handlers.AuthHandler{Handler: route.Handler}
	userHandler := handlers.UserHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	authenticationRoutes := route.RouteGroup.Group("/auth")

	authenticationRoutes.Post("/register", authHandler.Register)
	authenticationRoutes.Post("/login", authHandler.Login)

	authenticationRoutesJWT := authenticationRoutes.Group("/user")
	authenticationRoutesJWT.Use(jwtMiddleware.New)

	authenticationRoutesJWT.Get("", userHandler.Browse)
	authenticationRoutesJWT.Put("/:id", userHandler.Edit)
	authenticationRoutesJWT.Delete("/:id", userHandler.Delete)
}
