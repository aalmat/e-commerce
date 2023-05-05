package handler

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		}
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

func (h *Handler) signIn(ctx *gin.Context) {
	var input models.SignUser
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"token": token})

}
