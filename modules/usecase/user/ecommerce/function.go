package ecommerce

import (
	"fmt"
	"math"

	ee "basic-coding-kulina/modules/entity/ecommerce"
	ep "basic-coding-kulina/modules/entity/product"
)

func (ec *ecommerceUseCase) GetProductEcommerce(products *[]ep.Product, offset, pageSize int) (*[]ee.ProductResponse, int64, error) {
	var productResponses []ee.ProductResponse
	var productImage ep.ProductImage

	products, count, err := ec.ecommerceRepo.GetAllProduct(products, offset, pageSize)
	if err != nil {
		return &productResponses, count, err
	}

	for _, product := range *products {
		productImages, err := ec.ecommerceRepo.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return &productResponses, count, err
		}

		var productImageURLs []string
		for _, image := range productImages {
			productImageURLs = append(productImageURLs, image.ProductImageUrl)
		}

		exist, reviews, err := ec.ecommerceRepo.GetProductByID(product.ID)
		if err != nil {
			return &productResponses, count, err
		}

		var avgRating float64
		if exist {
			avgRating, err = ec.ecommerceRepo.AvgRating(product.ID)
			if err != nil {
				return &productResponses, count, err
			}
		} else {
			avgRating = 0
		}

		productResponse := ee.ProductResponse{
			ProductId:       product.ID,
			Name:            product.Name,
			Category:        product.ProductCategory.Category,
			Stock:           int(product.Stock),
			Weight:          product.Weight,
			Price:           product.Price,
			Status:          product.Status,
			Description:     product.Description,
			ProductImageUrl: productImageURLs,
			AvgRating:       math.Round(avgRating*10) / 10,
			Review:          reviews,
		}

		productResponses = append(productResponses, productResponse)
	}

	return &productResponses, count, nil
}
