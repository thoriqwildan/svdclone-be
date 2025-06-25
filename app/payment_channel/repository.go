package paymentchannel

import (
	"github.com/thoriqwildan/svdclone-be/pkg/database"
)

func GetFiltered(filter PaymentChannelFilter) ([]PaymentChannelResponse, int64, error) {
	var paymentChannels []PaymentChannelResponse
	var total int64

	query := database.DB.
		Table("payment_channels pc").
		Select(`
			pc.id,
			pc.name,
			pc.code,
			pc.icon_url,
			pc.order_num,
			pc.lib_name,
			pc.mdr,
			pc.fixed_fee,
			pc.created_at,
			pc.updated_at,
			pm.code as payment_method__code,
			pm.id as payment_method__id
		`).Joins("LEFT JOIN payment_methods pm ON pm.id = pc.payment_method_id")

	// Apply filters if provided
	if filter.Code != "" {
		query = query.Where("pc.code ILIKE ?", "%"+filter.Code+"%")
	}
	if filter.Name != "" {
		query = query.Where("pc.name ILIKE ?", "%"+filter.Name+"%")
	}

	// Get total count before applying pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}
	offset := (filter.Page - 1) * filter.Limit

	// Execute query and scan into paymentChannels slice
	if err := query.
		Limit(filter.Limit).
		Offset(offset).
		Order("pc.id DESC").
		Scan(&paymentChannels).Error; err != nil {
		return nil, 0, err
	}

	return paymentChannels, total, nil
}