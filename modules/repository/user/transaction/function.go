package transaction

import (
	"errors"

	ep "basic-coding-kulina/modules/entity/product"
	et "basic-coding-kulina/modules/entity/transaction"
	eu "basic-coding-kulina/modules/entity/user"
	ue "basic-coding-kulina/modules/entity/user"
)

func (tr *transactionRepo) GetUserById(id uint) (*ue.User, error) {
	user := &ue.User{}
	err := tr.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}

func (tr *transactionRepo) CreateTransaction(transaction *et.Transaction) error {
	err := tr.db.Create(&transaction).Error
	if err != nil {
		return err
	}

	//update stock
	for _, val := range transaction.TransactionDetails {
		var product ep.Product
		err := tr.db.Select("stock").Where("product_id = ?", val.ProductId).First(&product).Error
		if err != nil {
			return err
		}

		stock := product.Stock - val.Qty
		if stock == 0 {
			err = tr.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Updates(ep.Product{Stock: stock, Status: "habis"}).Error
		} else {
			err = tr.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Update("stock", stock).Error
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (tr *transactionRepo) UpdateTransaction(updateData et.Transaction) error {
	result := tr.db.Model(&et.Transaction{}).Where("transaction_id = ?", updateData.TransactionId).Updates(&updateData)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (tr *transactionRepo) GetPoint(id uint) (uint, error) {
	var userDetail eu.UserDetail

	if err := tr.db.Where("user_id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}
	point := userDetail.Point

	return point, nil

}
func (tr *transactionRepo) GetPaymentStatus(id string) (string, error) {
	var transaction et.Transaction

	if err := tr.db.Where("transaction_id = ?", id).First(&transaction).Error; err != nil {
		return "", err
	}
	status := transaction.PaymentStatus

	return status, nil

}

func (tr *transactionRepo) GetStock(id string) (uint, error) {
	var product ep.Product

	if err := tr.db.Where("product_id = ?", id).First(&product).Error; err != nil {
		return 0, err
	}
	stock := product.Stock

	return stock, nil

}

func (tr *transactionRepo) UpdatePoint(id uint, point uint) error {

	err := tr.db.Model(&eu.UserDetail{}).Where("user_id = ?", id).Update("point", point).Error
	if err != nil {
		return err
	}

	return nil
}
