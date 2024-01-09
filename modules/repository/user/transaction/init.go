package transaction

import (
	et "basic-coding-kulina/modules/entity/transaction"
	ue "basic-coding-kulina/modules/entity/user"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transaction *et.Transaction) error
	GetPoint(id string) (uint, error)
	GetStock(id string) (uint, error)
	UpdateTransaction(transaction et.Transaction) error
	GetPaymentStatus(id string) (string, error)
	GetUserById(id string) (*ue.User, error)
	UpdatePoint(id string, point uint) error
}

type transactionRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TransactionRepo {
	return &transactionRepo{
		db,
	}
}
