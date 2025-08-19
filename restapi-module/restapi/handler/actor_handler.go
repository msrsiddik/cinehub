package handler

import (
	"entities-module/model"
	"entities-module/query"
	"restapi-module/restapi/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ActorHandler interface {
	GetAllActors(c *fiber.Ctx) error
	GetActorByID(c *fiber.Ctx) error
}

type actorHandler struct {
	Q *query.Query
}

// NewActorHandler creates a new instance of ActorHandler
func NewActorHandler(query *query.Query) ActorHandler {
	return &actorHandler{
		Q: query,
	}
}

// GetAllActors godoc
// @Summary Get all actors
// @Description Returns a list of all actors
// @Tags Actors
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of actors per page" default(10)
// @Produce json
// @Success 200 {object} response.Paginated[model.Actor]
// @Router /actors [get]
func (h *actorHandler) GetAllActors(c *fiber.Ctx) error {
	// default pagination
	page, limit := 1, 10

	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	offset := (page - 1) * limit

	qActor := h.Q.Actor
	actors, err := qActor.WithContext(c.Context()).Limit(limit).Offset(offset).Find()
	if err != nil {
		return response.Fail(c, fiber.StatusInternalServerError, "Failed to fetch actors", err.Error())
	}

	totalCount, err := qActor.WithContext(c.Context()).Count()
	if err != nil {
		return response.Fail(c, fiber.StatusInternalServerError, "Failed to count actors", err.Error())
	}

	return response.Success(c, "Actors fetched successfully", &response.Paginated[*model.Actor]{
		Data:       actors,
		Page:       page, // Assuming page 1 for simplicity; implement pagination as needed
		Limit:      limit,
		Total:      len(actors),
		TotalPages: int(totalCount) / limit, // Calculate total pages
	})
}

// GetActorByID godoc
// @Summary Get actor by ID
// @Description Returns an actor by their ID
// @Tags Actors
// @Param id path int true "Actor ID"
// @Success 200 {object} model.Actor
// @Router /actors/{id} [get]
func (h *actorHandler) GetActorByID(c *fiber.Ctx) error {
	// This is a placeholder implementation.
	// Replace with actual logic to fetch an actor by ID from the database.
	idstr := c.Params("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid actor ID",
		})
	}
	// Convert int to int32 for the query
	id32 := int32(id)
	qActor := h.Q.Actor
	actor, err := qActor.WithContext(c.Context()).Where(qActor.ActorID.Eq(id32)).First()
	if err != nil {
		return response.Fail(c, fiber.StatusNotFound, "Actor not found", err.Error())
	}
	return response.Success(c, "Actor fetched successfully", actor)
}
