package seed

import (
	productEntity "basic-coding-kulina/modules/entity/product"
	"time"
)

func CreateProductImage() []*productEntity.ProductImage {
	product_images := []*productEntity.ProductImage{
		{
			ID:              "c476ce0c-84b7-4459-946b-e4c40b368c5c",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ID:              "8947ac5f-a101-4511-ad69-f579aecf7949",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ID:              "03bb3965-7b48-40cc-a09c-f7f85a1c642d",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ID:              "e764db51-0a9b-4798-a722-80cbff1c6e1b",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ID:              "fa66395e-4246-40df-b428-cce988685449",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "ba08266c-0926-484e-a610-c2a48ac6179d",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ID:              "7016de50-47ef-499b-93e9-24c81a5195b3",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "ba08266c-0926-484e-a610-c2a48ac6179d",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ID:              "c3b5b792-68c9-4cc5-9449-260df59269c0",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProductId:       "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
	}

	return product_images
}
