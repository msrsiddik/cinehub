package handler

import (
	"entities-module/model"
	"entities-module/query"
	"restapi-module/restapi/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddressHandler interface {
	GetAllAddresses(c *fiber.Ctx) error
	GetAddressByID(c *fiber.Ctx) error
}

type addressHandler struct {
	Q *query.Query
}

// NewAddressHandler creates a new instance of AddressHandler
func NewAddressHandler(query *query.Query) AddressHandler {
	return &addressHandler{
		Q: query,
	}
}

// GetAllAddresses godoc
// @Summary Get all addresses
// @Description Returns a list of all addresses
// @Tags Addresses
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of addresses per page" default(10)
// @Produce json
// @Success 200 {object} response.Paginated[model.Address]
// @Router /addresses [get]
func (h *addressHandler) GetAllAddresses(c *fiber.Ctx) error {
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
	hAddress := h.Q.Address
	addresses, err := hAddress.WithContext(c.Context()).Limit(limit).Offset(offset).Find()
	if err != nil {
		return response.Fail(c, fiber.StatusInternalServerError, "Failed to fetch addresses", err.Error())
	}

	totalCount, err := hAddress.WithContext(c.Context()).Count()
	if err != nil {
		return response.Fail(c, fiber.StatusInternalServerError, "Failed to count addresses", err.Error())
	}

	return response.Success(c, "Addresses fetched successfully", &response.Paginated[*model.Address]{
		Data:       addresses,
		Page:       page,
		Limit:      limit,
		Total:      len(addresses),
		TotalPages: int(totalCount) / limit,
	})
}

// GetAddressByID godoc
// @Summary Get address by ID
// @Description Returns an address by its ID
// @Tags Addresses
// @Param id path int true "Address ID"
// @Success 200 {object} model.Address
// @Router /addresses/{id} [get]
func (h *addressHandler) GetAddressByID(c *fiber.Ctx) error {
	idstr := c.Params("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid address ID",
		})
	}
	id32 := int32(id) // Convert int to int32 for the query

	qAddress := h.Q.Address
	address, err := qAddress.WithContext(c.Context()).Where(qAddress.AddressID.Eq(id32)).First()
	if err != nil {
		return response.Fail(c, fiber.StatusNotFound, "Address not found", err.Error())
	}
	return response.Success(c, "Address fetched successfully", address)
}
