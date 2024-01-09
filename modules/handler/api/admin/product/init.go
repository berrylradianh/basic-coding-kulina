package product

import (
	pc "basic-coding-kulina/modules/usecase/admin/product"
)

type ProductHandler struct {
	productUseCase pc.ProductUseCase
}

func New(productUseCase pc.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase,
	}
}
