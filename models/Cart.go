package models

type Cart struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
	ProductID uint `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID"`
	Quantity  uint `json:"quantity" binding:"required"`
}
