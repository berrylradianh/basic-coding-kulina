package transaction

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID            string `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	TransactionId string
	ProductId     string  `gorm:"size:255"  json:"ProductId" form:"ProductId" validate:"required"`
	ProductName   string  `json:"ProductName" form:"ProductName" validate:"required"`
	Qty           uint    `json:"Qty" form:"Qty" validate:"required"`
	SubTotalPrice float64 `json:"SubTotalPrice" form:"SubTotalPrice" validate:"required"`
}
