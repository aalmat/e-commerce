package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SearchByName(ctx *gin.Context) {

	search := ctx.Query("search")

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
		return
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

func (h *Handler) ViewSeller(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Product.ViewSeller(uint(uid))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func (h *Handler) FilterByPrice(ctx *gin.Context) {
	minPriceStr := ctx.Query("min_price")
	minPrice, err := strconv.ParseUint(minPriceStr, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	maxPriceStr := ctx.Query("max_price")
	maxPrice, err := strconv.ParseUint(maxPriceStr, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.FilterByPrice(uint(minPrice), uint(maxPrice))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": res,
	})

}
