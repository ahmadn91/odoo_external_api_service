package handlers

import (
	"github.com/ahmadn91/odoo_external_api_service/config"
	"github.com/ahmadn91/odoo_external_api_service/entities"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Value string
}

// fiber.Ctx is requests equivalent

func GetContacts(c *fiber.Ctx) error {
	ctx := context.Background()
	contacts, err := models.ResPartner.AllG(ctx)
	if err != nill {
		return c.JSON(err)
	}
	return c.Status(200).JSON(contacts)


}