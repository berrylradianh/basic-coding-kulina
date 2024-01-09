package product

import (
	pe "basic-coding-kulina/modules/entity/product"
	te "basic-coding-kulina/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (pr *productRepo) CreateProduct(product *pe.Product) error {
	if err := pr.db.Save(&product).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) CheckProductExist(productId string) (bool, error) {
	var count int64
	result := pr.db.Model(&pe.Product{}).Where("product_id = ?", productId).Count(&count)
	if result.Error != nil {
		return false, echo.NewHTTPError(500, result.Error)
	}

	exists := count > 0
	return exists, nil
}

func (pr *productRepo) CreateProductImage(productImage *pe.ProductImage) error {
	if err := pr.db.Save(&productImage).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) GetAllProduct(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	var count int64
	if err := pr.db.Model(&pe.Product{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := pr.db.
		Preload("ProductCategory").Preload("ProductImages").
		Offset(offset).
		Limit(pageSize).Order("created_at ASC").
		Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *products, count, nil
}

func (pr *productRepo) GetAllProductNoPagination(products *[]pe.Product) ([]pe.Product, error) {
	if err := pr.db.
		Preload("ProductCategory").
		Find(&products).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return *products, nil
}

func (pr *productRepo) GetProductByID(productId string, product *pe.Product) (*pe.Product, int64, float64, error) {
	var totalOrder int64
	var totalRevenue float64

	err := pr.db.Model(&te.Transaction{}).
		Select("COALESCE(SUM(transaction_details.qty),0) AS TotalOrder").
		Joins("JOIN transaction_details ON transactions.id = transaction_details.transaction_id").
		Joins("JOIN products ON products.id = transaction_details.product_id").
		Where("transactions.canceled_reason = ''").
		Where("transaction_details.product_id = ?", productId).
		Preload("ProductImages").
		Scan(&totalOrder).Error
	if err != nil {
		return nil, 0, 0, err
	}

	err = pr.db.Model(&te.Transaction{}).
		Select("COALESCE(SUM(transaction_details.sub_total_price),0) AS TotalRevenue").
		Joins("JOIN transaction_details ON transactions.id = transaction_details.transaction_id").
		Joins("JOIN products ON products.id = transaction_details.product_id").
		Where("transactions.canceled_reason = ''").
		Where("transaction_details.product_id = ?", productId).
		Preload("ProductImages").
		Scan(&totalRevenue).Error
	if err != nil {
		return nil, 0, 0, err
	}

	if err = pr.db.
		Preload("ProductCategory").Preload("ProductImages").
		Where("id = ?", productId).
		First(&product).Error; err != nil {
		return product, 0, 0, echo.NewHTTPError(404, err)
	}

	return product, totalOrder, totalRevenue, nil
}

func (pr *productRepo) UpdateProduct(productId string, req *pe.ProductRequest) error {
	if err := pr.db.Model(&pe.Product{}).Where("product_id = ?", productId).Updates(pe.Product{ProductCategoryId: req.ProductCategoryId, Name: req.Name, Price: req.Price, Status: req.Status, Description: req.Description}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) UpdateProductStock(productId string, stock uint) error {
	if err := pr.db.Exec("UPDATE products SET stock = ? WHERE product_id = ?", stock, productId).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) CountProductImage(productId string) (int, error) {
	var productImages *[]pe.ProductImage

	result := pr.db.Where("product_id = ?", productId).Find(&productImages)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (pr *productRepo) UpdateProductImage(idImage int, productId string, productImageUrl string) error {
	query := "UPDATE product_images SET product_image_url = ? WHERE product_id = ? AND product_image_url IN ( SELECT subquery.product_image_url FROM ( SELECT product_image_url FROM product_images WHERE product_id = ? LIMIT 1 OFFSET ? ) AS subquery)"
	if err := pr.db.Exec(query, productImageUrl, productId, productId, idImage-1).
		Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) DeleteProduct(productId string, product *pe.Product) error {
	// Hapus terlebih dahulu gambar produk yang terkait
	if err := pr.db.Where("product_id = ?", productId).Delete(&pe.ProductImage{}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	// Hapus produk dari tabel products
	if err := pr.db.Where("product_id = ?", productId).Delete(&product).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) SearchProduct(search, filter string, offset, pageSize int) (*[]pe.Product, int64, error) {
	var products []pe.Product
	var count int64

	if err := pr.db.Model(&pe.Product{}).
		Where("name LIKE ? OR id LIKE ? OR product_category_id IN (?)",
			"%"+search+"%",
			"%"+search+"%",
			pr.db.Model(&pe.ProductCategory{}).Select("id").Where("category LIKE ?", "%"+search+"%")).
		Where("status LIKE ?", "%"+filter+"%").
		Preload("ProductCategory").Preload("ProductImages").
		Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := pr.db.Model(&pe.Product{}).
		Where("name LIKE ? OR id LIKE ? OR product_category_id IN (?)",
			"%"+search+"%",
			"%"+search+"%",
			pr.db.Model(&pe.ProductCategory{}).Select("id").Where("category LIKE ?", "%"+search+"%")).
		Where("status LIKE ?", "%"+filter+"%").
		Preload("ProductCategory").Preload("ProductImages").
		Order("created_at ASC").
		Offset(offset).Limit(pageSize).
		Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return &products, count, nil
}
