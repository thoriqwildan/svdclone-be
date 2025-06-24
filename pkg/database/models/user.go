package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        	uint 
	Name				string
	Email				string `gorm:"unique"`
	ProfileUrl	sql.NullString
	Password		string
	Admin				bool `gorm:"default:false"`
	CreatedAt  	time.Time
	UpdatedAt  	time.Time 
}