package handlers

import (
	"context"
	"github.com/ahmadn91/odoo_external_api_service/models"
	"github.com/gofiber/fiber/v2"
	

)

type ErrorResponse struct {
	Value string
}

// fiber.Ctx is requests equivalent

func GetContacts(c *fiber.Ctx) error {
	ctx := context.Background()
	contacts, err := models.ResPartners().AllG(ctx)
	if err != nil {
		return c.JSON(err)
	}
	return c.Status(200).JSON(contacts)


}