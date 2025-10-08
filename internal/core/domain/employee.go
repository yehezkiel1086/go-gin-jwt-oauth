package domain

import "gorm.io/gorm"

type Employee struct {
	*gorm.Model

	Name string `json:"name" gorm:"size:255;unique"`
	Position string `json:"position" gorm:"size:255"`
}
