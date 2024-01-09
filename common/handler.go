package common

import (
	dh "basic-coding-kulina/modules/handler/api/admin/dashboard"
	oha "basic-coding-kulina/modules/handler/api/admin/order"
	ph "basic-coding-kulina/modules/handler/api/admin/product"
	pch "basic-coding-kulina/modules/handler/api/admin/product_category"
	ah "basic-coding-kulina/modules/handler/api/auth"
	ecommerceHandler "basic-coding-kulina/modules/handler/api/user/ecommerce"
	ohu "basic-coding-kulina/modules/handler/api/user/order"

	profileHandler "basic-coding-kulina/modules/handler/api/user/profile"
	uth "basic-coding-kulina/modules/handler/api/user/transaction"
)

type Handler struct {
	ProfileHandler         *profileHandler.ProfileHandler
	AuthHandler            *ah.AuthHandler
	TransactionHandlerUser *uth.TransactionHandler
	OrderHandlerUser       *ohu.OrderHandler
	ProductCategoryHandler *pch.ProductCategoryHandler
	ProductHandler         *ph.ProductHandler
	OrderHandlerAdmin      *oha.OrderHandlerAdmin
	DashboardHandler       *dh.DashboardHandler
	EcommerceHandler       *ecommerceHandler.EcommerceHandler
}
