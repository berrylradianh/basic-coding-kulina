package product

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID              string `gorm:"type:text;primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	ProductId       string         `gorm:"size:255" json:"-"`
	ProductImageUrl string         `json:"ProductImageUrl" form:"ProductImageUrl"`
	Product         Product        `json:"-"`
}
