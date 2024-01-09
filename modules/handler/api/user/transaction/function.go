package transaction

import (
	"net/http"

	h "basic-coding-kulina/helper/getIdUser"
	em "basic-coding-kulina/modules/entity/midtrans"
	er "basic-coding-kulina/modules/entity/rajaongkir"
	et "basic-coding-kulina/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := h.GetIdUser(c)

		transaction := et.Transaction{}
		c.Bind(&transaction)
		transaction.UserId = uint(id)

		snapUrl, transactionId, err := th.transactionUsecase.CreateTransaction(&transaction)
		if err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":         201,
			"Message":        "Success Create Transaction",
			"Transaction_Id": transactionId,
			"Payment_url":    snapUrl,
		})
	}

}

func (th *TransactionHandler) MidtransNotifications() echo.HandlerFunc {
	return func(c echo.Context) error {

		request := em.MidtransRequest{}
		c.Bind(&request)

		err := th.transactionUsecase.MidtransNotifications(&request)
		if err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Confirm Payment",
		})
	}

}

func (th *TransactionHandler) GetPoint() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := h.GetIdUser(c)

		res, err := th.transactionUsecase.GetPoint(id)

		if err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Get Point",
			"Point":   res,
		})
	}

}
func (th *TransactionHandler) GetPaymentStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")

		res, err := th.transactionUsecase.GetPaymentStatus(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":         200,
			"Message":        "Success Get Payment Status",
			"Payment Status": res,
		})
	}

}

func (th *TransactionHandler) ShippingOptions() echo.HandlerFunc {
	return func(c echo.Context) error {

		ship := er.RajaongkirRequest{}
		c.Bind(&ship)

		res, err := th.transactionUsecase.ShippingOptions(&ship)

		if err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Get Shipping Options",
			"Options": res,
		})
	}
}
