package models

type WareHouse struct {
	ID        uint `json:"ID"`
	ProductId uint `json:"product_Id" gorm:"ForeignKey:Product.ID"`
	Quantity  uint `json:"quantity"`
	UserID    uint `json:"userId" gorm:"ForeignKey:User.ID"`
	Price     int  `json:"price" binding:"required"`
}
