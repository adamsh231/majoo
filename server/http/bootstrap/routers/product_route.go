package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type ProductRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ProductRoutes) RegisterRoute() {
	//productHandler := handlers.AuthHandler{Handler: route.Handler}
	//jwtMiddleware := middlewares.JwtMiddleware{Contract:route.Handler.UcContract}
	//
	//productRoutes := route.RouteGroup.Group("/auth")
	//productRoutes.Use(jwtMiddleware.New)


}

