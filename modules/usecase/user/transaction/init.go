package transaction

import (
	em "basic-coding-kulina/modules/entity/midtrans"
	er "basic-coding-kulina/modules/entity/rajaongkir"
	et "basic-coding-kulina/modules/entity/transaction"
	rt "basic-coding-kulina/modules/repository/user/transaction"
)

type TransactionUsecase interface {
	CreateTransaction(transaction *et.Transaction) (string, string, error)
	GetPoint(id uint) (interface{}, error)
	MidtransNotifications(midtransRequest *em.MidtransRequest) error
	GetPaymentStatus(id string) (string, error)
	ShippingOptions(ship *er.RajaongkirRequest) (interface{}, error)
}

type transactionUsecase struct {
	transactionRepo rt.TransactionRepo
}

func New(adminRepo rt.TransactionRepo) *transactionUsecase {
	return &transactionUsecase{
		adminRepo,
	}
}
