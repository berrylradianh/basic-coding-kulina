package order

import (
	te "basic-coding-kulina/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (or *orderRepo) CheckOrderExist(transactionId string) (bool, error) {
	var count int64
	result := or.db.Model(&te.Transaction{}).Where("transaction_id = ?", transactionId).Count(&count)
	if result.Error != nil {
		return false, echo.NewHTTPError(500, result.Error)
	}

	exists := count > 0
	return exists, nil
}

func (or *orderRepo) GetAllOrder(transactions *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error) {
	var count int64
	if err := or.db.Model(&te.Transaction{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := or.db.Model(&te.Transaction{}).
		Select("*,(SELECT COUNT(*) FROM transaction_details WHERE transaction_details.transaction_id = transactions.id) AS Unit").
		Joins("JOIN transaction_details ON transaction_details.transaction_id = transactions.id").
		Joins("JOIN users ON transactions.user_id = users.id").
		Joins("JOIN user_details ON users.id = user_details.user_id").
		Offset(offset).
		Limit(pageSize).
		Scan(&transactions).Error; err != nil {
		return nil, 0, nil
	}

	return *transactions, count, nil
}

func (or *orderRepo) GetOrderByID(transactionId string, transaction *te.TransactionDetailResponse) (te.TransactionDetailResponse, error) {
	if err := or.db.Model(&te.Transaction{}).
		Select("*").
		Joins("JOIN users ON  users.id = transactions.user_id").
		Joins("JOIN user_details ON users.id = user_details.user_id").
		Joins("JOIN user_addresses ON transactions.address_id = user_addresses.id").
		Where("transaction_id = ?", transactionId).
		Scan(&transaction).Error; err != nil {
		return *transaction, err
	}
	return *transaction, nil
}

func (or *orderRepo) GetOrderProducts(transactionId string, products *[]te.TransactionProductDetailResponse) ([]te.TransactionProductDetailResponse, error) {
	var transaction te.Transaction
	if err := or.db.Model(&te.Transaction{}).Where("transaction_id = ?", transactionId).First(&transaction).Error; err != nil {
		return nil, err
	}

	if err := or.db.Model(&te.TransactionDetail{}).
		Select("*").
		Joins("JOIN products ON products.id = transaction_details.product_id").
		Where("transaction_details.transaction_id = ?", transaction.ID).Scan(&products).Error; err != nil {
		return nil, err
	}
	return *products, nil
}

func (or *orderRepo) SearchOrder(search, filter string, offset, pageSize int) (*[]te.TransactionResponse, int64, error) {
	var transactions []te.TransactionResponse
	var count int64

	if err := or.db.Model(&te.Transaction{}).
		Where("receipt_number LIKE ? OR transactions.transaction_id LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		).
		Where("status_transaction LIKE ?", "%"+filter+"%").
		Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := or.db.Model(&te.Transaction{}).
		Select("transactions.receipt_number AS ReceiptNumber, transactions.transaction_id AS TransactionId, user_details.name AS Name, (SELECT COUNT(*) FROM transaction_details WHERE transaction_details.transaction_id = transactions.id) AS Unit, total_price AS TotalPrice, transactions.created_at AS OrderDate, status_transaction AS StatusTransaction").
		Joins("JOIN transaction_details ON transaction_details.transaction_id = transactions.id").
		Joins("JOIN users ON transactions.user_id = users.id").
		Joins("JOIN user_details ON users.id = user_details.user_id").
		Where("receipt_number LIKE ? OR transactions.transaction_id LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		).
		Where("status_transaction LIKE ?", "%"+filter+"%").
		Offset(offset).Limit(pageSize).Find(&transactions).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return &transactions, count, nil
}

func (or *orderRepo) UpdateReceiptNumber(transactionId string, receiptNumber string) error {
	if err := or.db.Model(&te.Transaction{}).Where("transaction_id = ? AND status_transaction LIKE \"Dikemas\"", transactionId).Updates(te.Transaction{ReceiptNumber: receiptNumber, StatusTransaction: "Dikirim"}).Error; err != nil {
		return err
	}

	return nil
}
