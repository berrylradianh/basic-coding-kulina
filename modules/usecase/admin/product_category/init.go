package product_category

import (
	pe "basic-coding-kulina/modules/entity/product"
	pcr "basic-coding-kulina/modules/repository/admin/product_category"
)

type ProductCategoryUsecase interface {
	CreateProductCategory(productCategory *pe.ProductCategory) (bool, error)
	UpdateProductCategory(productCategory *pe.ProductCategory, id int) (bool, error)
	DeleteProductCategory(productCategory *pe.ProductCategory, id int) error
	GetAllProductCategory(offset, pageSize int) (*[]pe.ProductCategory, int64, error)
	GetProductCategoryById(id int) (*pe.ProductCategory, error)
	GetAllProductCategoryNoPagination() (*[]pe.ProductCategory, error)
	SearchingProductCategoryByName(name string, offset, pageSize int) (*[]pe.ProductCategory, int64, error)
}

type productCategoryUsecase struct {
	productCategoryRepo pcr.ProductCategoryRepo
}

func New(productCategoryRepo pcr.ProductCategoryRepo) *productCategoryUsecase {
	return &productCategoryUsecase{
		productCategoryRepo,
	}
}
