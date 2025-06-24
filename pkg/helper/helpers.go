package helper

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid: s != "",
	}
}

func ToNullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(i),
		Valid: i != 0,
	}
}