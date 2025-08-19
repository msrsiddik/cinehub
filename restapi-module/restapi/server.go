package restapi

import (
	_ "entities-module/model" // Import the database package to ensure it initializes
	"entities-module/query"
	_ "restapi-module/docs" // This is required to load the docs
	"restapi-module/restapi/handler"
	"restapi-module/restapi/router"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Cinehub REST API
// @version 1.0
// @description REST API for Cinehub using Fiber and GORM
// @host localhost:3000
// @BasePath /api/v1
// @schemes http https
// @contact.name Cinehub Support
// @contact.email msrsiddik2@gmail.com

func RestApiServer(app *fiber.App, query *query.Query) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/healthz", GetHealthz)

	actorHandler := handler.NewActorHandler(query)
	router.NewActorRouter(apiV1, actorHandler)

	addressHandler := handler.NewAddressHandler(query)
	router.NewAddressRouter(apiV1, addressHandler)
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns status ok
// @Tags Health
// @Success 200 {object} map[string]string
// @Router /healthz [get]
func GetHealthz(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
