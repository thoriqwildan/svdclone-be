package paymentmethod

import "time"

type CreatePaymentMethodRequest struct {
	Name				string `json:"name" validate:"required"`
	Desc				string `json:"desc,omitempty"`
	OrderNum		int    `json:"order_num,omitempty" validate:"omitempty,numeric"`
	UserAction	string `json:"user_action" validate:"required"`
	Code				string `json:"code,omitempty" validate:"required"`
}

type PaymentMethodResponse struct {
	Id 				uint   `json:"id"`
	Name 			string `json:"name"`
	Desc 			string `json:"desc,omitempty"`
	OrderNum	int    `json:"order_num,omitempty"`
	UserAction	string `json:"user_action"`
	Code 			string `json:"code,omitempty"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}

type PaymentMethodFilter struct {
	Code string `query:"code"`
	Name string `query:"name"`
}