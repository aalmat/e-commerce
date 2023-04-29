package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddToCart(ctx *gin.Context) {
	//userId, _ := h.GetUserId(ctx)
	userRole, _ := h.GetUserRole(ctx)

	if userRole != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, "you are not client")
		return
	}

}
