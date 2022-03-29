package api

import (
	"fmt"
	"time"

	"github.com/Shreeyash-Naik/Valet-Parking/models"
	"github.com/gofiber/fiber/v2"
)

// ../api/checkin
func Checkin(ctx *fiber.Ctx) error {
	body := new(CheckinBody)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	vehicle := models.Vehicle{
		CheckinTime:   time.Now().Format(time.RFC822),
		VehicleNumber: body.VehicleNumber,
		VehicleType:   body.VehicleType,
		VehicleModel:  body.VehicleModel,
		OwnerPhone:    body.OwnerPhone,
	}

	if err := Db.Model(&models.Vehicle{}).Create(&vehicle).Error; err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(vehicle)
}

// ../api/checkout
func Checkout(ctx *fiber.Ctx) error {
	query := new(CheckoutQuery)
	if err := ctx.QueryParser(query); err != nil {
		return err
	}

	var vehicle models.Vehicle
	if err := Db.Where("ID = ?", query.ID).Find(&vehicle).Error; err != nil {
		return err
	}

	checkinTime, _ := time.Parse(time.RFC822, vehicle.CheckinTime)
	hoursElasped := time.Since(checkinTime).Hours()

	if err := Db.Where("ID = ?", query.ID).Delete(&vehicle).Error; err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":           "Deleted",
		"hours_elapsed": hoursElasped,
		"total_cost":    fmt.Sprintf("%.2f", (hoursElasped)*COST_PER_HOUR),
	})
}

// ../api/search
func Search(ctx *fiber.Ctx) error {
	query := new(SearchQuery)
	if err := ctx.QueryParser(query); err != nil {
		return fiber.ErrBadRequest
	}

	query_ := fmt.Sprintf("%%%s%%", query.Query)

	var vehicles []models.Vehicle
	if err := Db.Where("vehicle_number LIKE ?", query_).Or("vehicle_model LIKE ?", query_).Find(&vehicles).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(vehicles)
}
