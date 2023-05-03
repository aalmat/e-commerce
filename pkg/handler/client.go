package handler

import (
	"errors"
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

	var quantity models.ProductQuantity
	if err := ctx.BindJSON(&quantity); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cartId, err := h.service.Client.AddToCart(userId, uint(uid), quantity.Quantity)
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

	ctx.JSON(http.StatusOK, wh)
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

	var quantity models.ProductQuantity
	if err := ctx.BindJSON(&quantity); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cartId, err := h.service.Client.ChangeProductQuantity(userId, uint(uid), quantity.Quantity)
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
		logrus.Println(errors.New("user role doesn't exists"))
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

	var comment models.Commentary
	if err := ctx.BindJSON(&comment); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment.ProductId = uint(productId)
	comment.UserId = userId
	commentId, err := h.service.Client.WriteComment(comment)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
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

	var rate models.Rating
	if err := ctx.BindJSON(&rate); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rate.ProductId = uint(productId)
	rate.UserId = userId
	rateId, err := h.service.Client.RateProduct(rate)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": rateId,
	})

}

func (h *Handler) PurchaseAll(ctx *gin.Context) {
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not client")
		return
	}

	err = h.service.Client.PurchaseAll(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "purchase")
}

func (h *Handler) PurchaseById(ctx *gin.Context) {
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

	err = h.service.Client.PurchaseById(userId, uint(uid))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"order id": id,
	})

}

func (h *Handler) ViewOrders(ctx *gin.Context) {
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

	if role != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not client")
		return
	}

	orders, err := h.service.Client.ShowOrders(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
