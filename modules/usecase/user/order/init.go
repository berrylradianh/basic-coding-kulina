package order

import (
	eo "basic-coding-kulina/modules/entity/order"
	ro "basic-coding-kulina/modules/repository/user/order"
)

type OderUsecase interface {
	GetOrder(filter string, idUser string, offset int, pageSize int) (interface{}, int64, error)
	Tracking(resi string, courier string) (interface{}, error)
	ConfirmOrder(eo.ConfirmOrder) error
	CancelOrder(co eo.CanceledOrder) error
}

type orderUsecase struct {
	orderRepo ro.OrderRepo
}

func New(orderRepo ro.OrderRepo) *orderUsecase {
	return &orderUsecase{
		orderRepo,
	}
}
