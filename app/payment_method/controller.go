package paymentmethod

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/thoriqwildan/svdclone-be/pkg/database"
	"github.com/thoriqwildan/svdclone-be/pkg/database/models"
	"github.com/thoriqwildan/svdclone-be/pkg/global"
	"github.com/thoriqwildan/svdclone-be/pkg/helper"
)

func CreatePaymentMethod(c *fiber.Ctx) error {
	var req CreatePaymentMethodRequest

	c.BodyParser(&req)
	if err := helper.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors: helper.TranslateErrorMessage(err),
		})
	}

	var paymentMethod models.PaymentMethod
	paymentMethod.Name = req.Name
	paymentMethod.Desc = helper.ToNullString(req.Desc)
	paymentMethod.OrderNum = req.OrderNum
	paymentMethod.UserAction = req.UserAction
	paymentMethod.Code = helper.ToNullString(req.Code)

	if err := database.DB.Create(&paymentMethod).Error; err != nil {
		log.Error("Failed to create payment method:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to create payment method",
			Errors:  nil,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment method created successfully",
		Data: PaymentMethodResponse{
			Id: paymentMethod.ID,
			Name: paymentMethod.Name,
			Desc: paymentMethod.Desc.String,
			OrderNum: paymentMethod.OrderNum,
			UserAction: paymentMethod.UserAction,
			Code: paymentMethod.Code.String,
			CreatedAt: paymentMethod.CreatedAt,
			UpdatedAt: paymentMethod.UpdatedAt,
		},
	})
}

func GetPaymentMethods(c *fiber.Ctx) error {
	var filter PaymentMethodFilter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Invalid query parameters",
			Errors:  err.Error(),
		})
	}

	data, _, err := GetFiltered(filter)
	if err != nil {
		log.Error("Failed to get payment methods:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to get payment methods",
			Errors:  nil,
		})
	}

	return c.JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment methods retrieved successfully",
		Data: data,
	})
}

func GetPaymentMethodById(c *fiber.Ctx) error {
	id := c.Params("id")

	var paymentMethod models.PaymentMethod
	if err := database.DB.First(&paymentMethod, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method not found",
			Errors:  nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment method retrieved successfully",
		Data: PaymentMethodResponse{
			Id: paymentMethod.ID,
			Name: paymentMethod.Name,
			Desc: paymentMethod.Desc.String,
			OrderNum: paymentMethod.OrderNum,
			UserAction: paymentMethod.UserAction,
			Code: paymentMethod.Code.String,
			CreatedAt: paymentMethod.CreatedAt,
			UpdatedAt: paymentMethod.UpdatedAt,
		},
	})
}

func UpdatePaymentMethod(c *fiber.Ctx) error {
	var req CreatePaymentMethodRequest
	id := c.Params("id")
	c.BodyParser(&req)
	if err := helper.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors: helper.TranslateErrorMessage(err),
		})
	}

	var paymentMethod models.PaymentMethod
	if err := database.DB.First(&paymentMethod, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment Method not Found!",
			Errors: nil,
		})
	}

	if err := database.DB.Where("name = ?", req.Name).First(&paymentMethod).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method with the same name already exists",
			Errors:  nil,
		})
	}

	paymentMethod.Name = req.Name
	paymentMethod.Desc = helper.ToNullString(req.Desc)
	paymentMethod.OrderNum = req.OrderNum
	paymentMethod.UserAction = req.UserAction
	paymentMethod.Code = helper.ToNullString(req.Code)
	if err := database.DB.Save(&paymentMethod).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to update payment method",
			Errors:  nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment method updated successfully",
		Data: PaymentMethodResponse{
			Id: paymentMethod.ID,
			Name: paymentMethod.Name,
			Desc: paymentMethod.Desc.String,
			OrderNum: paymentMethod.OrderNum,
			UserAction: paymentMethod.UserAction,
			Code: paymentMethod.Code.String,
			CreatedAt: paymentMethod.CreatedAt,
			UpdatedAt: paymentMethod.UpdatedAt,
		},
	})
}

func DeletePaymentMethod(c *fiber.Ctx) error {
	id := c.Params("id")

	var paymentMethod models.PaymentMethod
	if err := database.DB.First(&paymentMethod, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method not found",
			Errors:  nil,
		})
	}

	if err := database.DB.Delete(&paymentMethod).Error; err != nil {
		log.Error("Failed to delete payment method:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to delete payment method",
			Errors:  nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment method deleted successfully",
		Data: nil,
	})
}