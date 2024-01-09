package seed

import (
	pct "basic-coding-kulina/modules/entity/product"
)

func CreateProductCategory() *[]pct.ProductCategory {
	return &[]pct.ProductCategory{
		{
			Category: "perabot",
		},
		{
			Category: "kantong",
		},
		{
			Category: "Lainya",
		},
	}
}
