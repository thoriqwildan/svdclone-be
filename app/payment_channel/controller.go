package paymentchannel

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	paymentmethod "github.com/thoriqwildan/svdclone-be/app/payment_method"
	"github.com/thoriqwildan/svdclone-be/pkg/database"
	"github.com/thoriqwildan/svdclone-be/pkg/database/models"
	"github.com/thoriqwildan/svdclone-be/pkg/global"
	"github.com/thoriqwildan/svdclone-be/pkg/helper"
)

func CreatePaymentChannel(c *fiber.Ctx) error {
	var req PaymentChannelRequest
	var paymentChannel models.PaymentChannel

	c.BodyParser(&req)
	if err := helper.Validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helper.TranslateErrorMessage(err),
		})
	}

	// Check if payment method exists
	if err := database.DB.Where("id = ?", req.PaymentMethodId).First(&models.PaymentMethod{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method not found",
			Errors:  nil,
		})
	}

	// Check for duplicate payment channel code or name
	if err := database.DB.Where("code = ?", req.Code).Or("name = ?", req.Name).First(&paymentChannel).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment channel with the same code or name already exists",
			Errors:  nil,
		})
	}

	// Populate payment channel model from request
	paymentChannel.Name = req.Name
	paymentChannel.PaymentMethodId = req.PaymentMethodId
	paymentChannel.Code = req.Code
	paymentChannel.IconUrl = helper.ToNullString(req.IconUrl)
	paymentChannel.OrderNum = helper.ToNullInt64(req.OrderNum)
	paymentChannel.LibName = helper.ToNullString(req.LibName)
	paymentChannel.MDR = strconv.Itoa(req.Mdr)
	paymentChannel.FixedFee = req.FixedFee
	paymentChannel.UserAction = req.UserAction

	// Create payment channel in database
	if err := database.DB.Create(&paymentChannel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to create payment channel",
			Errors:  nil,
		})
	}

	// Respond with success and the created payment channel data
	return c.Status(fiber.StatusCreated).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel created successfully",
		Data: PaymentChannelResponse{
			Id:          paymentChannel.ID,
			Name:        paymentChannel.Name,
			Code:        paymentChannel.Code,
			IconUrl:     paymentChannel.IconUrl.String,
			OrderNum:    int(paymentChannel.OrderNum.Int64),
			LibName:     paymentChannel.LibName.String,
			Mdr:         paymentChannel.MDR,
			FixedFee:    paymentChannel.FixedFee,
			CreatedAt:   paymentChannel.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   paymentChannel.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			PaymentMethod: &PaymentMethod{ // Instantiate as a pointer
				Id:   paymentChannel.PaymentMethodId,
				Code: paymentmethod.GetCodeById(int(paymentChannel.PaymentMethodId)), // Assuming GetCodeById returns the code
			},
		},
	})
}

// GetPaymentChannels retrieves a list of payment channels based on filters.
func GetPaymentChannels(c *fiber.Ctx) error {
	var filter PaymentChannelFilter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Invalid query parameters",
			Errors:  err.Error(),
		})
	}

	data, total, err := GetFiltered(filter)
	if err != nil {
		log.Error("Failed to get payment methods:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to get payment methods",
			Errors:  nil,
		})
	}

	// Respond with paginated payment channel data
	return c.JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment methods retrieved successfully",
		Data: global.PaginationData{
			Items: data,
			Meta: global.PaginationPage{
				Page:  filter.Page,
				Limit: filter.Limit,
				Total: total,
			},
		},
	})
}

// GetPaymentChannelById retrieves a single payment channel by its ID.
func GetPaymentChannelById(c *fiber.Ctx) error {
	id := c.Params("id")
	var paymentChannel models.PaymentChannel
	if err := database.DB.Where("id = ?", id).First(&paymentChannel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment channel not found",
			Errors:  nil,
		})
	}

	// Respond with the retrieved payment channel data
	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel retrieved successfully",
		Data: PaymentChannelResponse{
			Id:          paymentChannel.ID,
			Name:        paymentChannel.Name,
			Code:        paymentChannel.Code,
			IconUrl:     paymentChannel.IconUrl.String,
			OrderNum:    int(paymentChannel.OrderNum.Int64),
			LibName:     paymentChannel.LibName.String,
			Mdr:         paymentChannel.MDR,
			FixedFee:    paymentChannel.FixedFee,
			CreatedAt:   paymentChannel.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   paymentChannel.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			PaymentMethod: &PaymentMethod{ // Instantiate as a pointer
				Id:   paymentChannel.PaymentMethodId,
				Code: paymentmethod.GetCodeById(int(paymentChannel.PaymentMethodId)), // Assuming GetCodeById returns the code
			},
		},
	})
}

// UpdatePaymentChannel handles updating an existing payment channel.
func UpdatePaymentChannel(c *fiber.Ctx) error {
	var req PaymentChannelRequest
	id := c.Params("id")

	c.BodyParser(&req)

	if err := helper.Validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helper.TranslateErrorMessage(err),
		})
	}

	var paymentChannel models.PaymentChannel
	if err := database.DB.First(&paymentChannel, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment channel not found",
			Errors:  nil,
		})
	}

	// Check if payment method exists
	if err := database.DB.Where("id = ?", req.PaymentMethodId).First(&models.PaymentMethod{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method not found",
			Errors:  nil,
		})
	}

	// Update payment channel model fields
	paymentChannel.Name = req.Name
	paymentChannel.PaymentMethodId = req.PaymentMethodId
	paymentChannel.Code = req.Code
	paymentChannel.IconUrl = helper.ToNullString(req.IconUrl)
	paymentChannel.OrderNum = helper.ToNullInt64(req.OrderNum)
	paymentChannel.LibName = helper.ToNullString(req.LibName)
	paymentChannel.MDR = strconv.Itoa(req.Mdr)
	paymentChannel.FixedFee = req.FixedFee
	paymentChannel.UserAction = req.UserAction

	// Save updated payment channel to database
	if err := database.DB.Save(&paymentChannel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to update payment channel",
			Errors:  nil,
		})
	}

	// Respond with success and the updated payment channel data
	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel updated successfully",
		Data: PaymentChannelResponse{
			Id:          paymentChannel.ID,
			Name:        paymentChannel.Name,
			Code:        paymentChannel.Code,
			IconUrl:     paymentChannel.IconUrl.String,
			OrderNum:    int(paymentChannel.OrderNum.Int64),
			LibName:     paymentChannel.LibName.String,
			Mdr:         paymentChannel.MDR,
			FixedFee:    paymentChannel.FixedFee,
			CreatedAt:   paymentChannel.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   paymentChannel.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			PaymentMethod: &PaymentMethod{ // Instantiate as a pointer
				Id:   paymentChannel.PaymentMethodId,
				Code: paymentmethod.GetCodeById(int(paymentChannel.PaymentMethodId)), // Assuming GetCodeById returns the code
			},
		},
	})
}

// DeletePaymentChannel handles the deletion of a payment channel by ID.
func DeletePaymentChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	var paymentChannel models.PaymentChannel
	if err := database.DB.First(&paymentChannel, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment channel not found",
			Errors:  nil,
		})
	}

	if err := database.DB.Delete(&paymentChannel).Error; err != nil {
		log.Error("Failed to delete payment channel:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to delete payment channel",
			Errors:  nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel deleted successfully",
		Data:    nil,
	})
}
