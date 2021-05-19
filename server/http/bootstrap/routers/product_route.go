package routers

import (
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/adamsh231/majoo/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ProductRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ProductRoutes) RegisterRoute() {
	productHandler := handlers.ProductHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract:route.Handler.UcContract}

	productRoutes := route.RouteGroup.Group("/product")
	productRoutes.Use(jwtMiddleware.New)

	productRoutes.Get("", productHandler.Browse)
	productRoutes.Get("/:id", productHandler.Read)
	productRoutes.Post("", productHandler.Add)
}

