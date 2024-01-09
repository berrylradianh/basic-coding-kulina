package product

import (
	"time"

	et "basic-coding-kulina/modules/entity/transaction"

	"gorm.io/gorm"
)

type Product struct {
	ID                 string `gorm:"type:text;primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt         `gorm:"index"`
	Name               string                 `validate:"required,max=10"`
	Stock              uint                   `validate:"required"`
	Price              float64                `validate:"required"`
	Status             string                 `validate:"required"`
	Weight             float64                `validate:"required"`
	Rating             float64                `validate:"required"`
	Description        string                 `validate:"required"`
	ProductCategoryId  string                 `json:"-" validate:"required"`
	ProductCategory    ProductCategory        `gorm:"foreignKey:ProductCategoryId" json:"-"`
	ProductImages      []ProductImage         `gorm:"foreignKey:ProductId"`
	TransactionDetails []et.TransactionDetail `gorm:"foreignKey:ProductId"`
}

type ProductRequest struct {
	ProductCategoryId string   `json:"productCategoryId" form:"productCategoryId"`
	Name              string   `json:"name" form:"name"`
	Stock             uint     `json:"stock" form:"stock"`
	Price             float64  `json:"price" form:"price"`
	Weight            float64  `validate:"required"`
	Description       string   `json:"description" form:"description"`
	Status            string   `json:"status" form:"status"`
	ProductImageUrl   []string `json:"productImageUrl" form:"productImageUrl"`
}

type ProductResponse struct {
	ID              string
	Name            string
	Category        string
	Stock           uint
	TotalOrders     uint    `json:"TotalOrders,omitempty"`
	TotalRevenue    float64 `json:"TotalRevenue,omitempty"`
	Weight          float64 `validate:"required"`
	Price           float64
	Status          string
	Rating          float64
	Description     string
	ProductImageUrl []string `json:"ProductImageUrl,omitempty"`
}
