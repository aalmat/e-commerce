package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
