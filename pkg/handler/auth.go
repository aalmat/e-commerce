package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//fmt.Printf("Password: %s", input.Password)

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})

}

func (h *Handler) signIn(stx *gin.Context) {

}
