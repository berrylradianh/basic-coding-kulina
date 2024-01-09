package app

import (
	"basic-coding-kulina/cmd/routes"
	"basic-coding-kulina/common"

	"basic-coding-kulina/database/psql"
	authHandler "basic-coding-kulina/modules/handler/api/auth"
	authRepo "basic-coding-kulina/modules/repository/auth"
	authUsecase "basic-coding-kulina/modules/usecase/auth"

	dashboardHandler "basic-coding-kulina/modules/handler/api/admin/dashboard"
	dashboardRepo "basic-coding-kulina/modules/repository/admin/dashboard"
	dashboardUsecase "basic-coding-kulina/modules/usecase/admin/dashboard"

	productCategoryHandler "basic-coding-kulina/modules/handler/api/admin/product_category"
	productCategoryRepo "basic-coding-kulina/modules/repository/admin/product_category"
	productCategoryUsecase "basic-coding-kulina/modules/usecase/admin/product_category"

	productHandler "basic-coding-kulina/modules/handler/api/admin/product"
	productRepo "basic-coding-kulina/modules/repository/admin/product"
	productUseCase "basic-coding-kulina/modules/usecase/admin/product"

	transactionHandlerUser "basic-coding-kulina/modules/handler/api/user/transaction"
	transactionRepoUser "basic-coding-kulina/modules/repository/user/transaction"
	transactionUsecaseUser "basic-coding-kulina/modules/usecase/user/transaction"

	orderHandlerUser "basic-coding-kulina/modules/handler/api/user/order"
	orderRepoUser "basic-coding-kulina/modules/repository/user/order"
	orderUsecaseUser "basic-coding-kulina/modules/usecase/user/order"

	orderHandlerAdmin "basic-coding-kulina/modules/handler/api/admin/order"
	orderRepoAdmin "basic-coding-kulina/modules/repository/admin/order"
	orderUsecaseAdmin "basic-coding-kulina/modules/usecase/admin/order"

	profileHandler "basic-coding-kulina/modules/handler/api/user/profile"
	profileRepo "basic-coding-kulina/modules/repository/user/profile"
	profileUsecase "basic-coding-kulina/modules/usecase/user/profile"

	ecommerceHandler "basic-coding-kulina/modules/handler/api/user/ecommerce"
	ecommerceRepo "basic-coding-kulina/modules/repository/user/ecommerce"
	ecommerceUseCase "basic-coding-kulina/modules/usecase/user/ecommerce"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	psql.Init()

	authRepo := authRepo.New(psql.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	productCategoryRepo := productCategoryRepo.New(psql.DB)
	productCategoryUsecase := productCategoryUsecase.New(productCategoryRepo)
	productCategoryHandler := productCategoryHandler.New(productCategoryUsecase)

	productRepo := productRepo.New(psql.DB)
	productUsecase := productUseCase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	ecommerceRepo := ecommerceRepo.New(psql.DB)
	ecommerceUsecase := ecommerceUseCase.New(ecommerceRepo)
	ecommerceHandler := ecommerceHandler.New(ecommerceUsecase)

	profileRepo := profileRepo.New(psql.DB)
	profileUsecase := profileUsecase.New(profileRepo)
	profileHandler := profileHandler.New(profileUsecase)

	transactionRepoUser := transactionRepoUser.New(psql.DB)
	transactionUsecaseUser := transactionUsecaseUser.New(transactionRepoUser)
	transactionHandlerUser := transactionHandlerUser.New(transactionUsecaseUser)

	orderRepoUser := orderRepoUser.New(psql.DB)
	orderUsecaseUser := orderUsecaseUser.New(orderRepoUser)
	orderHandlerUser := orderHandlerUser.New(orderUsecaseUser)

	orderRepoAdmin := orderRepoAdmin.New(psql.DB)
	orderUsecaseAdmin := orderUsecaseAdmin.New(orderRepoAdmin)
	orderHandlerAdmin := orderHandlerAdmin.New(orderUsecaseAdmin)

	dashboardRepo := dashboardRepo.New(psql.DB)
	dashboardUsecase := dashboardUsecase.New(dashboardRepo)
	dashboardHandler := dashboardHandler.New(dashboardUsecase)

	handler := common.Handler{
		ProductHandler:         productHandler,
		ProfileHandler:         profileHandler,
		AuthHandler:            authHandler,
		TransactionHandlerUser: transactionHandlerUser,
		OrderHandlerUser:       orderHandlerUser,
		ProductCategoryHandler: productCategoryHandler,
		OrderHandlerAdmin:      orderHandlerAdmin,
		DashboardHandler:       dashboardHandler,
		EcommerceHandler:       ecommerceHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
