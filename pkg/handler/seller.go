package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)

	var input models.WareHouse
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.ProductId = uint(uid)
	input.UserID = sellerId

	productId, err := h.service.Seller.AddProduct(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": productId,
	})

}

func (h *Handler) GetSellerProduct(ctx *gin.Context) {
	sellerId, err := h.GetUserId(ctx)
	//fmt.Println(sellerId)
	if err != nil {
		logrus.Println(err.Error())
		return
	}

	products, err := h.service.Seller.GetAllSellerProduct(sellerId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.GetAllWareHousesResponse{
		products,
	})
}

func (h *Handler) UpdateWareHouse(ctx *gin.Context) {
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

	var input models.UpdateWareHouse
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.WhId = uint(uid)
	err = h.service.Seller.UpdateProduct(sellerId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		"ok",
	})
}

func (h *Handler) DeleteWareHouse(ctx *gin.Context) {
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

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.Seller.DeleteProduct(sellerId, uint(uid))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": uid,
	})
}
