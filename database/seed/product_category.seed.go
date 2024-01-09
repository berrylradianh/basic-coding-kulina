package seed

import (
	pct "basic-coding-kulina/modules/entity/product"
	"time"
)

func CreateProductCategory() *[]pct.ProductCategory {
	return &[]pct.ProductCategory{
		{
			ID:        "ec716b5a-1a25-4096-ada0-403c9a24b914",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Category:  "perabot",
		},
		{
			ID:        "a2799195-042c-445a-878e-42e26df5e22f",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Category:  "kantong",
		},
		{
			ID:        "ee4f3a65-2cb5-4b46-b6e4-53f89207a8cc",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Category:  "Lainnya",
		},
	}
}
