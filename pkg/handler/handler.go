package handler

import (
	"github.com/aalmat/e-commerce/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const BasePage = "/market"

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *Handler) Routes() *gin.Engine {
	router := gin.New()

	ecommerce := router.Group(BasePage)
	auth := ecommerce.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
	}

	seller := ecommerce.Group("/seller", h.UserIdentify)
	{
		seller.GET("/", h.GetSellerProduct)
		seller.POST("/:id", h.AddProduct)
		seller.PUT("/:id", h.UpdateWareHouse)    //ware_id
		seller.DELETE("/:id", h.DeleteWareHouse) //ware_id
	}

	client := ecommerce.Group("/client", h.UserIdentify)
	{
		client.GET("/", h.ShowCartProducts)
		client.POST("/ware/:ware_id", h.AddToCart)            //ware
		client.DELETE("/cart/:cart_id", h.DeleteFromCart)     //cart
		client.PUT("/cart/:cart_id", h.ChangeProductQuantity) //cart
		client.POST("/:id/comment", h.WriteComment)           //product
		client.POST("/:id/rate", h.RateProduct)               //product
		client.POST("/purchase", h.PurchaseAll)
		client.POST("cart/:cart_id/purchase", h.PurchaseById) //cart
		client.GET("/orders", h.ViewOrders)
	}

	admin := ecommerce.Group("/admin", h.UserIdentify)
	{
		admin.POST("/", h.CreateProduct)
		admin.DELETE("/:id", h.DeleteProduct)
		admin.PUT("/:id", h.UpdateProduct)
		admin.GET("/", h.GetProducts)

	}

	product := ecommerce.Group("/product")
	{
		product.GET("/search", h.SearchByName)
		product.GET("/", h.GetProducts)
		product.GET("/:id", h.GetProductById)
		product.GET("/:id/comments", h.ViewComment)
		product.GET("/seller/:id", h.ViewSeller)
		product.GET("/filter", h.FilterByPrice) //products?min_price=10&max_price=50
	}

	return router
}
