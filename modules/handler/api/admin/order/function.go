package order

import (
	"math"
	"net/http"
	"strconv"

	te "basic-coding-kulina/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (oh *OrderHandlerAdmin) GetAllOrder(c echo.Context) error {
	var transactions []te.TransactionResponse

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	transactions, total, err := oh.orderUseCase.GetAllOrder(&transactions, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Gagal mendapatkan data pesanan",
			"Status":  http.StatusInternalServerError,
		})
	}

	if len(transactions) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list transaksi",
			"Status":  http.StatusNotFound,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Orders":    transactions,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

func (oh *OrderHandlerAdmin) GetOrderByID(c echo.Context) error {
	transactionId := c.Param("id")

	exist, err := oh.orderUseCase.CheckOrderExist(transactionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Message": "Pesanan tidak ditemukan",
			"Status":  http.StatusBadRequest,
		})
	}

	if exist {
		var transaction te.TransactionDetailResponse
		transaction, err := oh.orderUseCase.GetOrderByID(transactionId, &transaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan data pesanan",
				"Status":  http.StatusInternalServerError,
			})
		}

		var products []te.TransactionProductDetailResponse
		products, err = oh.orderUseCase.GetOrderProducts(transactionId, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan data produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		orderDetail := struct {
			Transaction te.TransactionDetailResponse
			Products    []te.TransactionProductDetailResponse
		}{
			Transaction: transaction,
			Products:    products,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Orders": orderDetail,
			"Status": http.StatusOK,
		})
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"Message": "Gagal mengambil detail pesanan",
		"Status":  http.StatusBadRequest,
	})
}

func (oh *OrderHandlerAdmin) SearchOrder(c echo.Context) error {
	var transactions *[]te.TransactionResponse

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	search := c.QueryParam("search")
	filter := c.QueryParam("filter")

	validParams := map[string]bool{"search": true, "filter": true, "page": true}

	for param := range c.QueryParams() {
		if !validParams[param] {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Masukkan parameter dengan benar",
				"Status":  http.StatusBadRequest,
			})
		}
	}

	transactions, total, err := oh.orderUseCase.SearchOrder(search, filter, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Gagal mendapatkan data pesanan",
			"Status":  http.StatusInternalServerError,
		})
	}

	if len(*transactions) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Pesanan yang anda cari tidak ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Orders":    transactions,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

func (oh *OrderHandlerAdmin) UpdateReceiptNumber(c echo.Context) error {
	transactionId := c.Param("id")
	receiptNumber := c.FormValue("ReceiptNumber")

	if receiptNumber == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Message": "Nomor resi tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	}

	exist, err := oh.orderUseCase.CheckOrderExist(transactionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Message": "Pesanan tidak ditemukan",
			"Status":  http.StatusBadRequest,
		})
	}

	if exist {
		var transaction te.TransactionDetailResponse
		transaction, err := oh.orderUseCase.GetOrderByID(transactionId, &transaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan data pesanan",
				"Status":  http.StatusInternalServerError,
			})
		}

		if transaction.StatusTransaction == "Dikemas" {
			err = oh.orderUseCase.UpdateReceiptNumber(transactionId, receiptNumber)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal mengubah nomor resi pesanan",
					"Status":  http.StatusInternalServerError,
				})
			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil mengubah nomor resi pesanan",
				"Status":  http.StatusOK,
			})
		}
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"Message": "Gagal mengubah nomor resi pesanan",
		"Status":  http.StatusBadRequest,
	})
}
