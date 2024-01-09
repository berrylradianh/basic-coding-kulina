package order

import (
	te "basic-coding-kulina/modules/entity/transaction"
	or "basic-coding-kulina/modules/repository/admin/order"
)

type OrderUseCase interface {
	CheckOrderExist(transactionId string) (bool, error)
	GetAllOrder(transaction *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error)
	GetOrderByID(transactionId string, transaction *te.TransactionDetailResponse) (te.TransactionDetailResponse, error)
	GetOrderProducts(transactionId string, products *[]te.TransactionProductDetailResponse) ([]te.TransactionProductDetailResponse, error)
	SearchOrder(search, filter string, offset, pageSize int) (*[]te.TransactionResponse, int64, error)
	UpdateReceiptNumber(transactionId string, receiptNumber string) error
}

type orderUseCase struct {
	orderRepo or.OrderRepo
}

func New(orderRepo or.OrderRepo) *orderUseCase {
	return &orderUseCase{
		orderRepo,
	}
}
