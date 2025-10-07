package domain

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Email string `json:"email" gorm:"size:255;unique"`
	Password string `json:"password" gorm:"size:255"`
}
