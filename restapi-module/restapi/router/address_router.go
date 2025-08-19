package router

import (
	"restapi-module/restapi/handler"

	"github.com/gofiber/fiber/v2"
)

func NewAddressRouter(grp fiber.Router, addressHandler handler.AddressHandler) fiber.Router {
	addressRoute := grp.Group("/addresses")
	addressRoute.Get("/", addressHandler.GetAllAddresses)   // Get all addresses
	addressRoute.Get("/:id", addressHandler.GetAddressByID) // Get address by ID
	return addressRoute
}
