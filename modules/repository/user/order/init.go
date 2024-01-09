package order

import (
	eo "basic-coding-kulina/modules/entity/order"

	"gorm.io/gorm"
)

type OrderRepo interface {
	GetOrder(filter string, idUser string, offset int, pageSize int) (interface{}, int64, error)
	// OrderDetail(id uint) (et.Transaction, error)
	// GetNameProductandImageUrl(id uint) (string, string, error)
	// GetPromoName(id uint) (string, error)
	ConfirmOrder(id string) error
	GetStatusOrder(id string) (string, error)
	CancelOrder(co eo.CanceledOrder) error
}

type orderRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db,
	}
}
