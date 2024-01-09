package transaction

import (
	"errors"
	"strconv"
	"time"

	"basic-coding-kulina/helper/hash"
	mdtrns "basic-coding-kulina/helper/midtrans"
	"basic-coding-kulina/helper/rajaongkir"
	vld "basic-coding-kulina/helper/validator"
	em "basic-coding-kulina/modules/entity/midtrans"
	er "basic-coding-kulina/modules/entity/rajaongkir"
	et "basic-coding-kulina/modules/entity/transaction"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (string, string, error) {

	user, err := tu.transactionRepo.GetUserById(transaction.UserId)
	if err != nil {
		return "", "", err
	}

	if user.RoleId == 1 {
		//lint:ignore ST1005 Reason for ignoring this linter
		return "", "", errors.New("Tidak boleh melakukan transaksi")
	}

	var productCost float64
	for _, cost := range transaction.TransactionDetails {
		stock, err := tu.transactionRepo.GetStock(cost.ProductId)
		if err != nil {
			return "", "", err
		}

		if stock < cost.Qty {
			//lint:ignore ST1005 Reason for ignoring this linter
			return "", "", errors.New("Qty melebihi stock")
		}
		productCost += cost.SubTotalPrice
	}

	transId := "eco" + strconv.FormatUint(uint64(transaction.UserId), 10) + time.Now().UTC().Format("2006010215040105")
	transaction.TransactionId = transId
	transaction.StatusTransaction = "Belum Bayar"
	transaction.TotalProductPrice = productCost
	transaction.TotalPrice = (transaction.TotalProductPrice + transaction.TotalShippingPrice) - (transaction.Point + transaction.Discount)

	redirectUrl, err := mdtrns.CreateMidtransUrl(transaction)
	if err != nil {
		return "", "", err
	}
	transaction.PaymentUrl = redirectUrl

	if err := vld.Validation(transaction); err != nil {
		return "", "", err
	}
	err = tu.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return "", "", err
	}

	if transaction.Point > 0 {
		res, err := tu.transactionRepo.GetPoint(transaction.UserId)
		if err != nil {
			return "", "", err
		}
		if transaction.Point > float64(res) {
			return "", "", errors.New("Tidak boleh melebihi point yang dimiliki")
		}
		point := res - uint(transaction.Point)
		err = tu.transactionRepo.UpdatePoint(transaction.UserId, point)
		if err != nil {
			return "", "", err
		}
	}

	return redirectUrl, transId, nil
}
func (tu *transactionUsecase) MidtransNotifications(midtransRequest *em.MidtransRequest) error {

	Key := hash.Hash(midtransRequest.OrderId, midtransRequest.StatusCode, midtransRequest.GrossAmount)
	if Key != midtransRequest.SignatureKey {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Invalid Transaction")
	}

	transaction := et.Transaction{
		//lint:ignore SA5011 Reason for ignoring this linter
		TransactionId: midtransRequest.OrderId,
		//lint:ignore SA5011 Reason for ignoring this linter
		PaymentStatus: midtransRequest.TransactionStatus,
		//lint:ignore SA5011 Reason for ignoring this linter
		PaymentMethod: midtransRequest.PaymentType,
	}

	if midtransRequest != nil {
		if midtransRequest.TransactionStatus == "capture" {
			if midtransRequest.FraudStatus == "challenge" {
				transaction.PaymentStatus = midtransRequest.FraudStatus
			} else if midtransRequest.FraudStatus == "accept" {
				transaction.StatusTransaction = "Dikemas"
				transaction.PaymentStatus = midtransRequest.FraudStatus
			}
		} else if midtransRequest.TransactionStatus == "settlement" {
			transaction.StatusTransaction = "Dikemas"
			transaction.PaymentStatus = midtransRequest.TransactionStatus
		} else if midtransRequest.TransactionStatus == "deny" {
			transaction.PaymentStatus = midtransRequest.TransactionStatus
		} else if midtransRequest.TransactionStatus == "cancel" || midtransRequest.TransactionStatus == "failure" {
			transaction.StatusTransaction = "Dibatalkan"
			transaction.CanceledReason = "Pembayaran gagal"
			transaction.PaymentStatus = midtransRequest.TransactionStatus
		} else if midtransRequest.TransactionStatus == "pending" {
			transaction.PaymentStatus = midtransRequest.TransactionStatus
		}
	}

	err := tu.transactionRepo.UpdateTransaction(transaction)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Invalid Transaction")
	}

	return nil
}
func (tu *transactionUsecase) GetPoint(id uint) (interface{}, error) {

	res, err := tu.transactionRepo.GetPoint(id)
	if err != nil {
		return 0, err
	}

	if res == 0 {
		return "Maaf, Kamu tidak punya point", nil
	}

	return res, nil
}
func (tu *transactionUsecase) GetPaymentStatus(id string) (string, error) {

	res, err := tu.transactionRepo.GetPaymentStatus(id)
	if err != nil {
		return "", err
	}
	return res, nil
}
func (tu *transactionUsecase) ShippingOptions(ship *er.RajaongkirRequest) (interface{}, error) {

	if err := vld.Validation(ship); err != nil {
		return nil, err
	}

	res, err := rajaongkir.ShippingOptions(ship)
	if err != nil {
		return nil, err
	}

	return res, nil

}
