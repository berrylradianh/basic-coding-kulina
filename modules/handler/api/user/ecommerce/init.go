package ecommerce

import (
	ec "basic-coding-kulina/modules/usecase/user/ecommerce"
)

type EcommerceHandler struct {
	ecommerceUseCase ec.EcommerceUsecase
}

func New(ecommerceUseCase ec.EcommerceUsecase) *EcommerceHandler {
	return &EcommerceHandler{
		ecommerceUseCase,
	}
}
