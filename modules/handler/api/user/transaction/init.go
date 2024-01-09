package transaction

import ut "basic-coding-kulina/modules/usecase/user/transaction"

type TransactionHandler struct {
	transactionUsecase ut.TransactionUsecase
}

func New(transactionUsecase ut.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase,
	}
}
