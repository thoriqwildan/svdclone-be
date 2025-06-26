package paymentmethod

import (
	"github.com/thoriqwildan/svdclone-be/pkg/database"
	"github.com/thoriqwildan/svdclone-be/pkg/database/models"
)

func GetFiltered(filter PaymentMethodFilter) ([]PaymentMethodResponse, int64, error) {
	var paymentMethods []PaymentMethodResponse
	var total int64

	query := database.DB.Model(&models.PaymentMethod{})

	if filter.Code != "" {
		query = query.Where("code ILIKE ?", "%"+filter.Code+"%")
	}
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Limit(filter.Limit).Offset(offset).Order("id DESC").Find(&paymentMethods).Error; err != nil {
		return nil, 0, err
	}

	return paymentMethods, total, nil
}

func GetCodeById(id int) string {
	var paymentMethod models.PaymentMethod

	if err := database.DB.Select("code").First(&paymentMethod, id).Error; err != nil {
		return ""
	}

	return paymentMethod.Code.String
}