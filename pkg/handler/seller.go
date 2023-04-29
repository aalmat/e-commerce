package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) AddProduct(ctx *gin.Context) {
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

	var input models.WareHouse
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.UserID = sellerId

	productId, err := h.service.AddProduct(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": productId,
	})

}
