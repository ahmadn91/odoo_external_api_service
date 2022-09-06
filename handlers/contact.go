package handlers

import (
	"context"
	"github.com/friendsofgo/errors"
	"fmt"
	"strconv"
	"github.com/lib/pq"
	"github.com/ahmadn91/odoo_external_api_service/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func GetContact(c *fiber.Ctx) error {
	ctx := context.Background()
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	contact, err := models.FindResPartnerG(ctx, int(id))
	if err != nil {
		return c.JSON(err)
	}
	return c.Status(200).JSON(contact)


}

func CreateContact(c *fiber.Ctx) error {
	ctx := context.Background()
	var contact models.ResPartner
	if err := c.BodyParser(contact); err != nil {
		return c.Status(503).SendString(err.Error()) 
	}

	err := contact.InsertG(ctx, boil.Infer() )
	if err != nil {
		err := errors.Cause(err).(*pq.Error)
        fmt.Printf("%+v\n", err.Code)
        return c.JSON(err)
    }
    return c.Status(201).JSON(contact)


}