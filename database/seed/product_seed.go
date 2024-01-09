package seed

import (
	"time"

	productEntity "basic-coding-kulina/modules/entity/product"
)

func CreateProduct() []*productEntity.Product {
	products := []*productEntity.Product{
		{
			ID:                "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			Name:              "Product Name 1",
			ProductCategoryId: "ec716b5a-1a25-4096-ada0-403c9a24b914",
			Stock:             10,
			Weight:            3000,
			Price:             30000,
			Status:            "tersedia",
			Description:       "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			ID:                "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductCategoryId: "a2799195-042c-445a-878e-42e26df5e22f",
			Name:              "Product Name 2",
			Stock:             100,
			Weight:            4000,
			Price:             36000,
			Status:            "tersedia",
			Rating:            0.00,
			Description:       "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			ID:                "ba08266c-0926-484e-a610-c2a48ac6179d",
			ProductCategoryId: "ee4f3a65-2cb5-4b46-b6e4-53f89207a8cc",
			Name:              "Product Name 3",
			Stock:             100,
			Weight:            5000,
			Price:             30000,
			Status:            "tersedia",
			Rating:            0.00,
			Description:       "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
	}

	return products
}
