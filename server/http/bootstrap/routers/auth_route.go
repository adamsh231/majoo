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
	AuthHandler := handlers.AuthHandler{Handler: route.Handler}
	UserHandler := handlers.UserHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract:route.Handler.UcContract}

	authenticationRoutes := route.RouteGroup.Group("/auth")

	authenticationRoutes.Post("/register", AuthHandler.Register)
	authenticationRoutes.Post("/login", AuthHandler.Login)

	authenticationRoutesJWT := authenticationRoutes.Group("/user")
	authenticationRoutesJWT.Use(jwtMiddleware.New)

	authenticationRoutesJWT.Get("", UserHandler.Browse)
	authenticationRoutesJWT.Put("/:id", UserHandler.Edit)
	authenticationRoutesJWT.Delete("/:id", UserHandler.Delete)
}

