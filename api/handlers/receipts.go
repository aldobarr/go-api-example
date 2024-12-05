package handlers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"regexp"
	"strings"
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

	purchaseDateTime, err := time.Parse(fmt.Sprintf("%s %s", time.DateOnly, "15:04"), fmt.Sprintf("%s %s", receipt.PurchaseDate, receipt.PurchaseTime))
	if err != nil {
		return HandleError(err, c, fiber.StatusUnprocessableEntity)
	}

	receipt.Points = 0
	alphaNumRetailer := regexp.MustCompile(`[^\p{L}\p{N}]+`).ReplaceAllString(receipt.Retailer, "")
	receipt.Points += len(alphaNumRetailer)

	if receipt.Total == math.Floor(receipt.Total) {
		receipt.Points += 50
	}

	if math.Mod(receipt.Total, 0.25) == 0 {
		receipt.Points += 25
	}

	itemCount := len(receipt.Items)
	itemPointsMultiple := int(math.Floor(float64(itemCount) / 2.0))
	receipt.Points += 5 * itemPointsMultiple

	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if math.Mod(float64(descLength), 3) == 0 {
			receipt.Points += int(math.Ceil(item.Price * 0.2))
		}
	}

	if math.Mod(float64(purchaseDateTime.Day()), 2) != 0 {
		receipt.Points += 6
	}

	// Requirements seem to specify between 2pm and 4pm exclusive.
	if purchaseDateTime.Hour() >= 14 && purchaseDateTime.Hour() < 16 && (purchaseDateTime.Hour() != 14 || purchaseDateTime.Minute() != 0) {
		receipt.Points += 10
	}

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

	data, err := database.Get(id)
	if err != nil {
		return HandleError(err, c, fiber.StatusInternalServerError)
	}

	receipt := new(Receipt)
	dec := gob.NewDecoder(bytes.NewReader(data))
	dec.Decode(receipt)

	return c.JSON(ReceiptPoints{receipt.Points})
}
