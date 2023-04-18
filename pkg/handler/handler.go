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

	product := ecommerce.Group("/products")
	{
		product.GET("/", h.CreateProduct)
	}

	return router
}
