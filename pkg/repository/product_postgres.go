package repository

import (
	"fmt"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"strings"
)

type ProductPostgres struct {
	db *gorm.DB
}

func (p *ProductPostgres) FilterByPrice(min, max uint) ([]models.ProductResponse, error) {
	var products []models.Product
	if err := p.db.Table("products").Joins("inner join ware_houses on products.id=ware_houses.product_id").Where("price between ? and ? and ware_houses.deleted_at is null and products.deleted_at is null", min, max).Scan(&products).Error; err != nil {
		return nil, err
	}

	resp := make([]models.ProductResponse, 0)
	for i := range products {
		res, err := p.GetProductDetails(products[i])
		if err != nil {
			return nil, err
		}
		resp = append(resp, res)
	}

	return resp, nil

}

func (p *ProductPostgres) ViewComment(productId uint) ([]models.CommentUser, error) {
	var result []models.CommentUser

	if err := p.db.Table("users").Select("users.email, commentaries.text, commentaries.created_at").Joins("inner join commentaries on commentaries.user_id = users.id").Where("commentaries.product_id = ?", productId).Order("commentaries.created_at").Scan(&result).Error; err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}

func (p *ProductPostgres) GetAll() ([]models.ProductResponse, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	resp := make([]models.ProductResponse, 0)
	for i := range products {
		res, err := p.GetProductDetails(products[i])
		if err != nil {
			return nil, err
		}
		resp = append(resp, res)
	}

	return resp, nil
}

func (p *ProductPostgres) GetById(productId uint) (models.ProductResponse, error) {
	var product models.Product
	if err := p.db.Where("id = ?", productId).First(&product).Error; err != nil {
		return models.ProductResponse{}, err
	}
	return p.GetProductDetails(product)
}

func (p *ProductPostgres) SearchByName(search string) ([]models.ProductResponse, error) {

	words := strings.Split(search, " ")
	query := ""
	for _, v := range words {
		if len(query) != 0 {
			query += " "
		}
		query += "%"
		query += v
		query += "%"
	}
	/*
		if err := c.db.Table("ware_houses").Select("distinct on (ware_house_id) carts.id as cart_id, ware_houses.product_id, ware_houses.user_id, ware_houses.price, ware_houses.created_at, ware_houses.id, carts.quantity").Joins("inner join carts on ware_houses.id = carts.ware_house_id").Where("carts.user_id=? and carts.deleted_at is null", userId).Scan(&whs).Error; err != nil {
				return nil, err
			}
	*/

	//fmt.Println(search)
	var products []models.Product
	if err := p.db.Where("title like ?", query).Find(&products).Error; err != nil {
		return nil, err
	}

	resp := make([]models.ProductResponse, 0)
	for i := range products {
		res, err := p.GetProductDetails(products[i])
		if err != nil {
			return nil, err
		}
		resp = append(resp, res)
	}

	return resp, nil
}

func (p *ProductPostgres) GetProductDetails(product models.Product) (models.ProductResponse, error) {

	var resp models.ProductResponse
	resp.Product = product
	if err := p.db.Where("product_id=?", product.ID).Find(&resp.WareHouses).Error; err != nil {
		return models.ProductResponse{}, err
	}
	var err error
	resp.Comments, err = p.ViewComment(product.ID)
	if err != nil {
		return models.ProductResponse{}, err
	}
	//fmt.Println(resp.Comments)

	return resp, nil
}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db}
}

func (p *ProductPostgres) ViewSeller(sellerId uint) ([]models.SellerResponse, error) {
	var res []models.SellerResponse
	if err := p.db.Table("products").Select("products.id as product_id, title, description, photo, rating, price").Joins("inner join ware_houses on products.id = ware_houses.product_id").Where("user_id = ? and products.deleted_at is null and ware_houses.deleted_at is null", sellerId).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
