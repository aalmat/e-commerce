package models

type SellerResponse struct {
	ProductId   uint   `json:"product_id"`
	Title       string `json:"title"`
	Description string `json:"description" `
	Photo       string `json:"photo"`
	Rating      uint   `json:"rating"`
	Price       uint   `json:"price"`
}
