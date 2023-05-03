package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SearchByName(ctx *gin.Context) {

	var search models.Search

	if err := ctx.BindJSON(&search); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.service.Product.SearchByName(search)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.GetAllListsResponse{Data: products})
}

func (h *Handler) GetProducts(ctx *gin.Context) {

	products, err := h.service.Product.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, models.GetAllListsResponse{
		Data: products,
	})
}

func (h *Handler) GetProductById(ctx *gin.Context) {

	id := ctx.Param("id")
	productId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	product, err := h.service.Product.GetById(uint(productId))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Handler) ViewComment(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Product.ViewComment(uint(uid))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)

}
