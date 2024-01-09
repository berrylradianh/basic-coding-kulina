package ecommerce

import (
	ee "basic-coding-kulina/modules/entity/ecommerce"
	ep "basic-coding-kulina/modules/entity/product"
	er "basic-coding-kulina/modules/repository/user/ecommerce"
)

type EcommerceUsecase interface {
	GetProductEcommerce(products *[]ep.Product, offset, pageSize int) (*[]ee.ProductResponse, int64, error)
}

type ecommerceUseCase struct {
	ecommerceRepo er.EcommerceRepo
}

func New(ecommerceRepo er.EcommerceRepo) *ecommerceUseCase {
	return &ecommerceUseCase{
		ecommerceRepo,
	}
}
