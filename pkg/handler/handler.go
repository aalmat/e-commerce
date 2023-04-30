package handler

import (
	"github.com/aalmat/e-commerce/pkg/service"
	"github.com/gin-gonic/gin"
)

const BasePage = "/market"

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
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
		seller.POST("/", h.AddProduct)
	}

	client := ecommerce.Group("/client", h.UserIdentify)
	{
		client.GET("")
		client.POST("/:id", h.AddToCart)
		client.DELETE("/:id", h.DeleteFromCart)
		client.DELETE("/:id", h.ChangeProductQuantity)
	}

	admin := ecommerce.Group("/admin", h.UserIdentify)
	{
		admin.POST("/", h.CreateProduct)
		admin.DELETE("/:id", h.DeleteProduct)
		admin.PUT("/:id", h.UpdateProduct)
		admin.GET("/", h.GetProducts)
	}

	return router
}
