package router

import (
	"restapi-module/restapi/handler"

	"github.com/gofiber/fiber/v2"
)

func NewActorRouter(grp fiber.Router, actorHandler handler.ActorHandler) fiber.Router {
	actorRoute := grp.Group("/actors")
	actorRoute.Get("/", actorHandler.GetAllActors)    // Get all actors
	actorRoute.Get("/:id", actorHandler.GetActorByID) // Get actor by ID
	return actorRoute
}
