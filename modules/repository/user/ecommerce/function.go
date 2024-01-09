package ecommerce

import (
	ee "basic-coding-kulina/modules/entity/ecommerce"
	ep "basic-coding-kulina/modules/entity/product"
)

func (er *ecommerceRepo) GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (er *ecommerceRepo) GetProductByID(productId string) (bool, []ee.ReviewResponse, error) {
	var reviewResponse *[]ee.ReviewResponse

	result := er.db.Raw("SELECT ud.name, ud.profile_photo, rp.rating, rp.comment, rp.comment_admin, rp.photo_url, rp.video_url FROM rating_products rp JOIN transaction_details td ON(rp.transaction_detail_id = td.id) JOIN transactions t ON(td.transaction_id = t.id) JOIN users u ON(t.user_id = u.id) JOIN user_details ud ON(u.id = ud.user_id) JOIN products p ON(td.product_id = p.product_id) WHERE p.product_id = ?", productId).Scan(&reviewResponse)
	if result.Error != nil {
		return false, nil, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil, result.Error

	}

	return true, *reviewResponse, nil
}

func (er *ecommerceRepo) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	var productImages []ep.ProductImage

	if err := er.db.Model(&ep.ProductImage{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, err
	}

	return productImages, nil
}

func (er *ecommerceRepo) AvgRating(productId string) (float64, error) {
	var avgRating float64

	if err := er.db.Raw("SELECT AVG(rp.rating) FROM rating_products rp JOIN transaction_details td ON(rp.transaction_detail_id = td.id) JOIN products p ON(td.product_id = p.product_id) WHERE p.product_id = ?", productId).Scan(&avgRating).Error; err != nil {
		return 0, err
	}

	return avgRating, nil
}
