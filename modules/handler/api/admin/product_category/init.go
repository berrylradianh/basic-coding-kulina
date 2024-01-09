package productcategory

import pcc "basic-coding-kulina/modules/usecase/admin/product_category"

type ProductCategoryHandler struct {
	productCategoryUsecase pcc.ProductCategoryUsecase
}

func New(productCategoryUsecase pcc.ProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		productCategoryUsecase,
	}
}
