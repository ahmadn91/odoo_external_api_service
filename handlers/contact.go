package handlers

import (
	"context"

	"github.com/ahmadn91/odoo_external_api_service/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Value string
}

var validate = validator.New()

func ValidateStruct(contact models.ResPartner) []*ErrorResponse {
    var errors []*ErrorResponse
    err := validate.Struct(contact)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element ErrorResponse
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
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