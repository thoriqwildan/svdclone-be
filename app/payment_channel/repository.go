package paymentchannel

import (
	"github.com/thoriqwildan/svdclone-be/pkg/database"
)

type PaymentChannelQueryResult struct {
	Id                uint    `json:"id"`
	Name              string  `json:"name"`
	Code              string  `json:"code"`
	IconUrl           string  `json:"icon_url"`
	OrderNum          int     `json:"order_num"`
	LibName           string  `json:"lib_name"`
	Mdr               string  `json:"mdr"`
	FixedFee          float64 `json:"fixed_fee"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	PaymentMethodId   uint    `json:"payment_method_id"`
	PaymentMethodCode string  `json:"payment_method_code"`
}

func GetFiltered(filter PaymentChannelFilter) ([]PaymentChannelResponse, int64, error) {
	var results []PaymentChannelQueryResult
	var total int64

	query := database.DB.
		Table("payment_channels pc").
		Select(`
			pc.id,
			pc.name,
			pc.code,
			COALESCE(pc.icon_url, '') as icon_url,
			COALESCE(pc.order_num, 0) as order_num,
			COALESCE(pc.lib_name, '') as lib_name,
			pc.mdr,
			pc.fixed_fee,
			TO_CHAR(pc.created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at,
			TO_CHAR(pc.updated_at, 'YYYY-MM-DD HH24:MI:SS') as updated_at,
			pc.payment_method_id,
			COALESCE(pm.code, '') as payment_method_code
		`).
		Joins("LEFT JOIN payment_methods pm ON pm.id = pc.payment_method_id")

	// Filter
	if filter.Code != "" {
		query = query.Where("pc.code ILIKE ?", "%"+filter.Code+"%")
	}
	if filter.Name != "" {
		query = query.Where("pc.name ILIKE ?", "%"+filter.Name+"%")
	}

	// Total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}
	offset := (filter.Page - 1) * filter.Limit

	// Query data
	if err := query.
		Limit(filter.Limit).
		Offset(offset).
		Order("pc.id DESC").
		Scan(&results).Error; err != nil {
		return nil, 0, err
	}

	// Convert to response format
	var responses []PaymentChannelResponse
	for _, result := range results {
		response := PaymentChannelResponse{
			Id:   result.Id,
			Name: result.Name,
			Code: result.Code,
			PaymentMethod: PaymentMethod{
				Id:   result.PaymentMethodId,
				Code: result.PaymentMethodCode,
			},
			IconUrl:   result.IconUrl,
			OrderNum:  result.OrderNum,
			LibName:   result.LibName,
			Mdr:       result.Mdr,
			FixedFee:  result.FixedFee,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, total, nil
}