package order

import (
	"errors"
	"strings"

	b "basic-coding-kulina/helper/binderbyte"
	vld "basic-coding-kulina/helper/validator"
	eo "basic-coding-kulina/modules/entity/order"
)

func (oc *orderUsecase) GetOrder(filter string, idUser uint, offset int, pageSize int) (interface{}, int64, error) {

	filter = strings.Trim(filter, "%20")
	res, count, err := oc.orderRepo.GetOrder(filter, idUser, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return "Belum ada pesanan", 0, nil
	}

	return res, count, nil
}

func (oc *orderUsecase) Tracking(resi string, courier string) (interface{}, error) {

	res, err := b.Tracking(resi, courier)
	if err != nil {
		return "", err
	}

	return res, nil
}
func (oc *orderUsecase) ConfirmOrder(confirmOrder eo.ConfirmOrder) error {

	if err := vld.Validation(confirmOrder); err != nil {
		return err
	}

	statusTransaction, err := oc.orderRepo.GetStatusOrder(confirmOrder.TransactionId)
	if err != nil {
		return err
	}

	if statusTransaction != "Dikirim" {
		return errors.New("Tidak bisa mengonfirmasi pesanan sebelum barang Dikirim")
	}

	err = oc.orderRepo.ConfirmOrder(confirmOrder.TransactionId)
	if err != nil {
		return err
	}

	return nil
}
func (oc *orderUsecase) CancelOrder(co eo.CanceledOrder) error {

	if err := vld.Validation(co); err != nil {
		return err
	}
	statusTransaction, err := oc.orderRepo.GetStatusOrder(co.TransactionId)
	if err != nil {
		return err
	}

	if statusTransaction != "Belum Bayar" {
		return errors.New("Tidak bisa membatalkan pesanan")
	}

	err = oc.orderRepo.CancelOrder(co)
	if err != nil {
		return err
	}

	return nil
}
