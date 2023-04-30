package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

/*
AddToCart(productId uint, quantity uint) (uint, error)
	ShowCartProducts() ([]models.Product, error)
	DeleteFromCart(productId uint, quantity uint) (uint, error)
	SearchByName(keyword string) ([]models.Product, error)
	FilterByPrice(minPrice, maxPrice int) ([]models.Product, error)
	FilterByRating(minRate, maxRate int) ([]models.Product, error)
*/

func (h *Handler) AddToCart(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not client")
		return
	}

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var quantity uint
	if err := ctx.BindJSON(&quantity); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cartId, err := h.service.Client.AddToCart(userId, uint(uid), quantity)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cart id": cartId,
	})
}

func (h *Handler) ShowCartProducts(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusBadRequest, "you are not client")
		return
	}

	wh, err := h.service.Client.ShowCartProducts(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.GetAllWareHousesResponse{wh})
}

func (h *Handler) DeleteFromCart(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusBadRequest, "you are not client")
		return
	}

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	h.service.Client.DeleteFromCart(userId, uint(uid))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": uid,
	})
}

func (h *Handler) ChangeProductQuantity(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusBadRequest, "you are not client")
		return
	}

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var quantity uint
	if err := ctx.BindJSON(&quantity); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cartId, err := h.service.Client.ChangeProductQuantity(userId, uint(uid), quantity)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cart id": cartId,
	})

}

func (h *Handler) WriteComment(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not client")
		return
	}

	id := ctx.Param("id")
	productId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var commentText string
	if err := ctx.BindJSON(&commentText); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commentId, err := h.service.Client.WriteComment(userId, uint(productId), commentText)
	if err := ctx.BindJSON(&commentText); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": commentId,
	})

}

func (h *Handler) RateProduct(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not client")
		return
	}

	id := ctx.Param("id")
	productId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var rate uint
	if err := ctx.BindJSON(&rate); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rateId, err := h.service.Client.RateProduct(userId, uint(productId), rate)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": rateId,
	})

}
