package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model `json:"-"`
	Code       string `gorm:"unique" json:"code"`
	Price      uint   `json:"price"`
}

func GetPro(code string) Product {
	var product Product
	if DB.First(&product, "code = ?", code).Error != nil {
		product.Code = "None"
	}
	return product
}
