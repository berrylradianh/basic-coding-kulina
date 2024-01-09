package ecommerce

import (
	ee "basic-coding-kulina/modules/entity/ecommerce"
	ep "basic-coding-kulina/modules/entity/product"

	"gorm.io/gorm"
)

type EcommerceRepo interface {
	GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	GetProductByID(productId string) (bool, []ee.ReviewResponse, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	AvgRating(productId string) (float64, error)
}

type ecommerceRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) EcommerceRepo {
	return &ecommerceRepo{
		db,
	}
}
