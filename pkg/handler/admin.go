package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProduct(ctx *gin.Context) {

	userRole, err := h.GetUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if userRole != models.Admin {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not admin")
		return
	}

	var input models.Product
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.validate.Struct(input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := h.service.Admin.CreateProduct(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": productId,
	})

}

func (h *Handler) UpdateProduct(ctx *gin.Context) {
	userRole, err := h.GetUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if userRole != models.Admin {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not admin")
		return
	}

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input models.ProductUpdate
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.validate.Struct(input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Admin.UpdateProduct(uint(uid), input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": uid,
	})

}

func (h *Handler) DeleteProduct(ctx *gin.Context) {
	userRole, err := h.GetUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if userRole != models.Admin {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not admin")
		return
	}

	id := ctx.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Admin.DeleteProduct(uint(uid))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": uid,
	})

}
