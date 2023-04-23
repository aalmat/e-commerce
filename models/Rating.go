package models

type Rating struct {
	Rate      uint `json:"rate"`
	UserId    uint `json:"userId" gorm:"ForeignKey:User.ID"`
	ProductId uint `json:"productId" gorm:"ForeignKey:Product.ID"`
}
