package paymentchannel

type PaymentChannelRequest struct {
	Name				string `json:"name" validate:"required"`
	Code 				string `json:"code,omitempty" validate:"required"`
	PaymentMethodId uint   `json:"payment_method_id" validate:"required"`
	IconUrl			string `json:"icon_url,omitempty"`
	OrderNum		int    `json:"order_num,omitempty" validate:"omitempty,numeric"`
	LibName			string `json:"lib_name,omitempty"`
	Mdr 				int `json:"mdr,omitempty" validate:"omitempty,numeric"`
	FixedFee		float64 `json:"fixed_fee,omitempty" validate:"omitempty,numeric"`
	UserAction	string `json:"user_action" validate:"required"`
}

type PaymentChannelResponse struct {
	Id                uint    `json:"id"`
	Name              string  `json:"name"`
	PaymentMethod 		PaymentMethod  `json:"payment_method"`
	Code              string  `json:"code,omitempty"`
	IconUrl           string  `json:"icon_url,omitempty"`
	OrderNum          int     `json:"order_num,omitempty"`
	LibName           string  `json:"lib_name,omitempty"`
	Mdr               string  `json:"mdr,omitempty"`
	FixedFee          float64 `json:"fixed_fee,omitempty"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type PaymentMethod struct {
	Id  uint   `json:"id"`
	Name string `json:"name"`
}

type PaymentChannelFilter struct {
	Code string `query:"code"`
	Name string `query:"name"`
	Page int    `query:"page"`
	Limit int    `query:"limit"`
}