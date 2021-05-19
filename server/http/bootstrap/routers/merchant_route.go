package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/adamsh231/majoo/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type MerchantRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route MerchantRoutes) RegisterRoute() {
	merchantHandler := handlers.MerchantHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract:route.Handler.UcContract}

	merchantRoutes := route.RouteGroup.Group("/merchant")
	merchantRoutes.Use(jwtMiddleware.New)

	merchantRoutes.Get("", merchantHandler.Browse)
	merchantRoutes.Post("", merchantHandler.Add)

}
