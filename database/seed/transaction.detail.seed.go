package seed

import (
	et "basic-coding-kulina/modules/entity/transaction"
	"time"
)

func CreateTransactionDetail() []*et.TransactionDetail {
	transaction := []*et.TransactionDetail{
		{
			ID:            "d13439db-6cde-41eb-bf7e-3c84686a9330",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			TransactionId: "1b25566d-f99c-4f2a-bd42-83c00e1b5a39",
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			ID:            "63fd6119-5533-43fe-af37-13771f522449",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			TransactionId: "c4cc586e-bade-4918-8205-abdead811e1a",
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			ID:            "3123737d-68d5-4f78-8d35-0fab6a515086",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			TransactionId: "1cad50aa-0f0e-4b85-b270-660a2c76fa96",
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			ID:            "95d0dff6-bec0-498d-8927-2d61ca6ea47f",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			TransactionId: "159c6856-b7d9-44b0-b5cc-476813a6b388",
			ProductId:     "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductName:   "Product Name 2",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			ID:            "20eae8b0-ddc5-4af0-ac71-f7d77e7ce690",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			TransactionId: "1829821b-19be-4957-a371-0ca5b494e1b6",
			ProductId:     "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductName:   "Product Name 2",
			Qty:           1,
			SubTotalPrice: 30000,
		},
	}
	return transaction
}
