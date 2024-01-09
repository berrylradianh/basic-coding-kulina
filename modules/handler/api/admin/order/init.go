package order

import (
	oc "basic-coding-kulina/modules/usecase/admin/order"
)

type OrderHandlerAdmin struct {
	orderUseCase oc.OrderUseCase
}

func New(orderUseCase oc.OrderUseCase) *OrderHandlerAdmin {
	return &OrderHandlerAdmin{
		orderUseCase,
	}
}
