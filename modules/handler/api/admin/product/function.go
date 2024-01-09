package product

import (
	"context"
	"fmt"
	"math"

	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"basic-coding-kulina/helper/cloudstorage"
	cs "basic-coding-kulina/helper/customstatus"

	vld "basic-coding-kulina/helper/validator"
	ep "basic-coding-kulina/modules/entity/product"

	"github.com/labstack/echo/v4"
)

func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	var products []ep.Product

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	products, total, err := h.productUseCase.GetAllProduct(&products, offset, pageSize)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list produk",
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

	var productResponses []ep.ProductResponse
	for _, product := range products {
		var imageURLs []string

		for _, image := range product.ProductImages {
			imageURLs = append(imageURLs, image.ProductImageUrl)
		}

		productResponse := ep.ProductResponse{
			ID:              product.ID,
			Name:            product.Name,
			Category:        product.ProductCategory.Category,
			Stock:           product.Stock,
			Weight:          product.Weight,
			Price:           product.Price,
			Status:          product.Status,
			Rating:          product.Rating,
			Description:     product.Description,
			ProductImageUrl: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Products":  productResponses,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	var product *ep.Product
	productID := c.Param("id")

	product, totalOrder, totalRevenue, err := h.productUseCase.GetProductByID(productID, product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	var imageURLs []string
	for _, image := range product.ProductImages {
		imageURLs = append(imageURLs, image.ProductImageUrl)
	}

	productResponse := ep.ProductResponse{
		ID:              product.ID,
		Name:            product.Name,
		Category:        product.ProductCategory.Category,
		Stock:           product.Stock,
		Weight:          product.Weight,
		TotalOrders:     uint(totalOrder),
		TotalRevenue:    totalRevenue,
		Price:           product.Price,
		Status:          product.Status,
		Rating:          product.Rating,
		Description:     product.Description,
		ProductImageUrl: imageURLs,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Products": productResponse,
		"Status":   http.StatusOK,
	})
}

func (h *ProductHandler) SearchProduct(c echo.Context) error {
	var products *[]ep.Product
	var err error

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

	products, total, err := h.productUseCase.SearchProduct(search, filter, offset, pageSize)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	if len(*products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Produk yang anda cari tidak ditemukan",
			"Status":  http.StatusNotFound,
		})
	} else {
		var productResponses []ep.ProductResponse
		for _, product := range *products {
			var imageURLs []string
			for _, image := range product.ProductImages {
				imageURLs = append(imageURLs, image.ProductImageUrl)
			}

			productResponse := ep.ProductResponse{
				ID:              product.ID,
				Name:            product.Name,
				Category:        product.ProductCategory.Category,
				Stock:           product.Stock,
				Weight:          product.Weight,
				Price:           product.Price,
				Status:          product.Status,
				Rating:          product.Rating,
				Description:     product.Description,
				ProductImageUrl: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Products":  productResponses,
			"Page":      page,
			"TotalPage": int(math.Ceil(float64(total) / float64(pageSize))),
			"Status":    http.StatusOK,
		})
	}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var product ep.Product
	productCategoryID := c.FormValue("ProductCategoryId")
	if productCategoryID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	} else {
		product.ID = productCategoryID
	}

	name := c.FormValue("Name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	} else {
		product.Name = name
	}

	weightStr := c.FormValue("Weight")
	if weightStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	} else {
		weight, err := strconv.ParseFloat(weightStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Data yang diisi harus angka",
				"Status":  http.StatusBadRequest,
			})
		}
		product.Weight = weight
	}

	stockStr := c.FormValue("Stock")
	if stockStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	} else {
		stock, err := strconv.ParseUint(stockStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Data yang diisi harus angka",
				"Status":  http.StatusBadRequest,
			})
		}
		product.Stock = uint(stock)
		if stock == 0 {
			product.Status = "habis"
		} else {
			product.Status = "tersedia"
		}
	}

	priceStr := c.FormValue("Price")
	if priceStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	} else {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Data yang diisi harus angka",
				"Status":  http.StatusBadRequest,
			})
		}
		product.Price = float64(price)
	}

	description := c.FormValue("Description")
	if description == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	} else {
		product.Description = description
	}

	err := h.productUseCase.CreateProduct(&product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	cloudstorage.Folder = "img/products/"
	photoUploaded := false
	for i := 1; i <= 5; i++ {
		fileHeader, err := c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			photoUploaded = true
			fileExtension := filepath.Ext(fileHeader.Filename)
			allowedExtensions := map[string]bool{
				".png":  true,
				".jpeg": true,
				".jpg":  true,
			}
			if !allowedExtensions[fileExtension] {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, format file yang Anda unggah tidak sesuai",
					"Status":  http.StatusBadRequest,
				})
			}
			maxFileSize := 4 * 1024 * 1024
			fileSize := fileHeader.Size
			if fileSize > int64(maxFileSize) {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB",
					"Status":  http.StatusBadRequest,
				})
			}

			PhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)

			productImage := ep.ProductImage{
				ID:              product.ID,
				ProductImageUrl: PhotoUrl,
			}
			err = h.productUseCase.CreateProductImage(&productImage)
			if err != nil {
				code, msg := cs.CustomStatus(err.Error())
				return c.JSON(code, echo.Map{
					"Status":  code,
					"Message": msg,
				})
			}

		} else {
			if err != nil {
				i = 1000
			}
		}
	}

	if !photoUploaded {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Field/gambar tidak boleh kosong",
			"Status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"Message": "Anda berhasil menambahkan produk",
		"Status":  http.StatusCreated,
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	productId := c.Param("id")
	var req ep.ProductRequest
	var product ep.Product

	productBefore, _, _, err := h.productUseCase.GetProductByID(productId, &product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	productCategoryID := c.FormValue("ProductCategoryId")
	if productCategoryID == "" {
		req.ProductCategoryId = productBefore.ProductCategoryId
	} else {
		req.ProductCategoryId = productCategoryID
	}

	weightStr := c.FormValue("Weight")
	if weightStr == "" {
		req.Weight = productBefore.Weight
	} else {
		weight, err := strconv.ParseFloat(weightStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Data yang diisi harus angka",
				"Status":  http.StatusBadRequest,
			})
		}
		req.Weight = weight
	}

	name := c.FormValue("Name")
	if name == "" {
		req.Name = productBefore.Name
	} else {
		req.Name = name
	}

	stockStr := c.FormValue("Stock")
	if stockStr == "" {
		req.Stock = productBefore.Stock
	} else {
		stock, err := strconv.ParseUint(stockStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Data yang diisi harus angka",
				"Status":  http.StatusBadRequest,
			})
		}
		req.Stock = uint(stock)
		if stock == 0 {
			req.Status = "habis"
		} else {
			req.Status = "tersedia"
		}
	}

	priceStr := c.FormValue("Price")
	if priceStr == "" {
		req.Price = productBefore.Price
	} else {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Data yang diisi harus angka",
				"Status":  http.StatusBadRequest,
			})
		}
		req.Price = float64(price)
	}

	description := c.FormValue("Description")
	if description == "" {
		req.Description = productBefore.Description
	} else {
		req.Description = description
	}

	err = h.productUseCase.UpdateProduct(productId, &req)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	err = h.productUseCase.UpdateProductStock(productId, req.Stock)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	nProductImage, err := h.productUseCase.CountProductImage(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal menghitung banyak foto tiap produk",
			"Status":  http.StatusInternalServerError,
		})
	}

	var fileHeader *multipart.FileHeader
	cloudstorage.Folder = "img/products/"

	for i := 1; i <= nProductImage; i++ {
		fileHeader, _ = c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			if productBefore.ProductImages[i-1].ProductImageUrl != "" {
				fileName := cloudstorage.GetFileName(productBefore.ProductImages[i-1].ProductImageUrl)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"Message": "Gagal mendapatkan nama file",
					})
				}
				err := cloudstorage.DeleteImage(fileName)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"Message": "Gagal menghapus file pada cloud storage",
					})
				}
			}

			if err := vld.ValidateFileExtension(fileHeader); err != nil {
				return err
			}

			maxFileSize := 4 * 1024 * 1024
			if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
				return err
			}

			photoUrl, err := cloudstorage.UploadToBucket(context.Background(), fileHeader)
			if err != nil {
				return err
			}

			if err := h.productUseCase.UpdateProductImage(i, productId, photoUrl); err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Status":  http.StatusInternalServerError,
					"Message": "fail update product image",
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil mengubah produk",
		"Status":  http.StatusOK,
	})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	productId := c.Param("id")

	var product *ep.Product
	product, _, _, err := h.productUseCase.GetProductByID(productId, product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	cloudstorage.Folder = "img/products/"
	for _, image := range product.ProductImages {
		filename := cloudstorage.GetFileName(image.ProductImageUrl)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mendapatkan nama file",
			})
		}
		err = cloudstorage.DeleteImage(filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal menghapus file pada cloud storage",
				"Status":  http.StatusInternalServerError,
			})
		}
	}

	err = h.productUseCase.DeleteProduct(productId, product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menghapus produk",
		"Status":  http.StatusOK,
	})
}

func (h *ProductHandler) DownloadCSVFile(c echo.Context) error {
	var products []ep.Product

	products, err := h.productUseCase.GetAllProductNoPagination(&products)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Berhasil membuat file CSV",
		"Status":  http.StatusOK,
		"Data":    products,
	})
}
