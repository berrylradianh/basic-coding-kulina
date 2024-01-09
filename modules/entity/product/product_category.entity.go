package product

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        string `gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Category  string         `json:"Category" form:"category" validate:"required"`
	Products  []Product      `gorm:"foreignKey:ProductCategoryId"`
}

type ProductCategoryResponse struct {
	Category string    `json:"category" form:"category"`
	Products []Product `gorm:"foreignKey:ProductCategoryId" json:"-"`
}
