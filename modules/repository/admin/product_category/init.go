package product_category

import (
	pe "basic-coding-kulina/modules/entity/product"

	"gorm.io/gorm"
)

type ProductCategoryRepo interface {
	CreateProductCategory(productCategory *pe.ProductCategory) error
	UpdateProductCategory(productCategory *pe.ProductCategory, id int) error
	DeleteProductCategory(productCategory *pe.ProductCategory, id int) error
	GetAllProductCategory(offset, pageSize int) (*[]pe.ProductCategory, int64, error)
	GetProductCategoryById(id int) (*pe.ProductCategory, error)
	GetAllProductCategoryNoPagination() (*[]pe.ProductCategory, error)
	SearchingProductCategoryByName(name string, offset, pageSize int) (*[]pe.ProductCategory, int64, error)
	IsProductCategoryAvailable(productCategory *pe.ProductCategory, name string) (bool, error)
}

type productCategoryRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProductCategoryRepo {
	return &productCategoryRepo{
		db,
	}
}
