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

	if err := database.DB.Where("id = ?", req.PaymentMethodId).First(&models.PaymentMethod{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method not found",
			Errors:  nil,
		})
	}

	if err := database.DB.Where("code = ?", req.Code).Or("name = ?", req.Name).First(&paymentChannel).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment channel with the same code or name already exists",
			Errors:  nil,
		})
	}

	paymentChannel.Name = req.Name
	paymentChannel.PaymentMethodId = req.PaymentMethodId
	paymentChannel.Code = req.Code
	paymentChannel.IconUrl = helper.ToNullString(req.IconUrl)
	paymentChannel.OrderNum = helper.ToNullInt64(req.OrderNum)
	paymentChannel.LibName = helper.ToNullString(req.LibName)
	paymentChannel.MDR = strconv.Itoa(req.Mdr)
	paymentChannel.FixedFee = req.FixedFee
	paymentChannel.UserAction = req.UserAction

	if err := database.DB.Create(&paymentChannel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to create payment channel",
			Errors:  nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel created successfully",
		Data:    PaymentChannelResponse{
			Id:	paymentChannel.ID,
			Name: paymentChannel.Name,
			PaymentMethod: PaymentMethod{
				Id: paymentChannel.PaymentMethodId,
				Code: paymentmethod.GetCodeById(int(paymentChannel.PaymentMethodId)),
			},
			Code: paymentChannel.Code,
			IconUrl: paymentChannel.IconUrl.String,
			OrderNum: int(paymentChannel.OrderNum.Int64),
			LibName: paymentChannel.LibName.String,
			Mdr: paymentChannel.MDR,
			FixedFee: paymentChannel.FixedFee,
		},
	})
}

func GetPaymentChannels(c *fiber.Ctx) error {
	var filter PaymentChannelFilter
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
	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel retrieved successfully",
		Data: PaymentChannelResponse{
			Id: paymentChannel.ID,
			Name: paymentChannel.Name,
			PaymentMethod: PaymentMethod{
				Id: paymentChannel.PaymentMethodId,
				Code: paymentmethod.GetCodeById(int(paymentChannel.PaymentMethodId)),
			},
			Code: paymentChannel.Code,
			IconUrl: paymentChannel.IconUrl.String,
			OrderNum: int(paymentChannel.OrderNum.Int64),
			LibName: paymentChannel.LibName.String,
			Mdr: paymentChannel.MDR,
			FixedFee: paymentChannel.FixedFee,
			CreatedAt: paymentChannel.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: paymentChannel.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

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

	if err := database.DB.Where("id = ?", req.PaymentMethodId).First(&models.PaymentMethod{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Payment method not found",
			Errors:  nil,
		})
	}

	// if err := database.DB.Where("code = ?", req.Code).Or("name = ?", req.Name).First(&paymentChannel).Error; err == nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
	// 		Success: false,
	// 		Message: "Payment channel with the same code or name already exists",
	// 		Errors:  nil,
	// 	})
	// }

	paymentChannel.Name = req.Name
	paymentChannel.PaymentMethodId = req.PaymentMethodId
	paymentChannel.Code = req.Code
	paymentChannel.IconUrl = helper.ToNullString(req.IconUrl)
	paymentChannel.OrderNum = helper.ToNullInt64(req.OrderNum)
	paymentChannel.LibName = helper.ToNullString(req.LibName)
	paymentChannel.MDR = strconv.Itoa(req.Mdr)
	paymentChannel.FixedFee = req.FixedFee
	paymentChannel.UserAction = req.UserAction
	if err := database.DB.Save(&paymentChannel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Failed to update payment channel",
			Errors:  nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(global.SuccessResponse{
		Success: true,
		Message: "Payment channel updated successfully",
		Data: PaymentChannelResponse{
			Id: paymentChannel.ID,
			Name: paymentChannel.Name,
			PaymentMethod: PaymentMethod{
				Id: paymentChannel.PaymentMethodId,
				Code: paymentmethod.GetCodeById(int(paymentChannel.PaymentMethodId)),
			},
			Code: paymentChannel.Code,
			IconUrl: paymentChannel.IconUrl.String,
			OrderNum: int(paymentChannel.OrderNum.Int64),
			LibName: paymentChannel.LibName.String,
			Mdr: paymentChannel.MDR,
			FixedFee: paymentChannel.FixedFee,
			CreatedAt: paymentChannel.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: paymentChannel.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

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