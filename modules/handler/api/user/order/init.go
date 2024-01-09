package order

import (
	uo "basic-coding-kulina/modules/usecase/user/order"
)

type OrderHandler struct {
	orderUsecase uo.OderUsecase
}

func New(informationUsecase uo.OderUsecase) *OrderHandler {
	return &OrderHandler{
		informationUsecase,
	}
}
