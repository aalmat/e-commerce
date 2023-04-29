package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateProduct(ctx *gin.Context) {
	sellerId, err := h.GetUserId(ctx)
	//fmt.Println(sellerId)
	if err != nil {
		logrus.Println(err.Error())
		return
	}
	role, err := h.GetUserRole(ctx)
	if err != nil {
		logrus.Println(err.Error())
		return
	}

	//fmt.Println(role)

	if role != models.Seller {
		newErrorResponse(ctx, http.StatusBadRequest, "you are not seller")
		return
	}

	var input models.Product
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := h.service.CreateProduct(sellerId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": productId,
	})

}


type GetAllListsResponse struct {
	Data []models.Product `json:"data"`
}

func (h *Handler) GetAll(ctx *gin.Context) {
	products, err := h.service.Product.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, GetAllListsResponse{
		products,
	})
}
