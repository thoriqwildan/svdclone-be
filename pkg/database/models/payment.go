package models

import (
	"database/sql"
	"time"
)

type PaymentMethod struct {
	ID 				uint
	Name 			string	`gorm:"unique"`
	Desc 			sql.NullString
	OrderNum	int	`gorm:"default:1"`
	UserAction	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Code 			sql.NullString

	Channels	[]PaymentChannel `gorm:"foreignKey:PaymentMethodId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PaymentChannel struct {
	ID 				uint
	PaymentMethodId	uint
	Code 			string	`gorm:"unique"`
	Name 			string `gorm:"unique"`
	IconUrl		sql.NullString
	OrderNum	sql.NullInt64	`gorm:"default:1"`
	LibName		sql.NullString
	UserAction	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	MDR       string  `gorm:"default:'0'"`
	FixedFee  float64 `gorm:"default:0"`
}