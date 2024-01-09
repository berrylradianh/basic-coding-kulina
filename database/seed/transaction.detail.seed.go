package seed

import (
	et "basic-coding-kulina/modules/entity/transaction"
)

func CreateTransactionDetail() []*et.TransactionDetail {
	transaction := []*et.TransactionDetail{
		{
			TransactionId: 1,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 2,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 3,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 4,
			ProductId:     "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductName:   "Product Name 2",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 5,
			ProductId:     "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductName:   "Product Name 2",
			Qty:           1,
			SubTotalPrice: 30000,
		},
	}
	return transaction
}
