package handlers

import (
	"fmt"
	"time"

	"github.com/aldobarr/go-api-example/api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ValidatePurchaseDate(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	if date == "" {
		return false
	}

	_, err := time.Parse(time.DateOnly, date)
	return err == nil
}

func ValidatePurchaseTime(fl validator.FieldLevel) bool {
	ptime := fl.Field().String()
	if ptime == "" {
		return false
	}

	_, err := time.Parse("15:04", ptime)
	return err == nil
}

func ProcessReceipt(c *fiber.Ctx) error {
	receipt := new(Receipt)
	err := c.BodyParser(receipt)
	if err != nil {
		return HandleError(err, c, fiber.StatusBadRequest)
	}

	verr := ValidateInput(receipt)
	if verr != nil {
		return HandleError(verr, c, fiber.StatusUnprocessableEntity)
	}

	receipt.Points = 0

	id := ReceiptID{uuid.New().String()}

	dberr := database.UpdateOrInsert(id.ID, receipt)

	if dberr != nil {
		return HandleError(dberr, c, fiber.StatusInternalServerError)
	}

	return c.JSON(id)
}

func GetPoints(c *fiber.Ctx) error {
	id := c.Params("id")

	if !database.Exists(id) {
		return HandleError(fmt.Errorf("Receipt with ID %s not found", id), c, fiber.StatusNotFound)
	}

	return c.SendString(id)
}
