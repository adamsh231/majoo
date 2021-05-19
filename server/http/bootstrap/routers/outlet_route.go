package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/adamsh231/majoo/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type OutletRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route OutletRoutes) RegisterRoute() {
	outletHandler := handlers.OutletHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract:route.Handler.UcContract}

	outletRoutes := route.RouteGroup.Group("/outlet")
	outletRoutes.Use(jwtMiddleware.New)

	outletRoutes.Get("", outletHandler.Browse)
	outletRoutes.Post("", outletHandler.Add)
}