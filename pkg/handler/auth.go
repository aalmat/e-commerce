package handler

import (
	"fmt"
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

type SignUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input SignUser
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

// store invalidated tokens in a map
var blacklist map[string]bool

// function to invalidate a token
func invalidateToken(token string) error {
	// check if token already exists in the blacklist
	if blacklist[token] {
		return fmt.Errorf("Token already invalidated")
	}
	// add token to blacklist
	blacklist[token] = true
	return nil
}

// handler for logout endpoint
//func logoutHandler(ctx *gin.Context) {
//	header := ctx.GetHeader(Authorization)
//	if header == "" {
//		newErrorResponse(ctx, http.StatusBadRequest, "header is empty")
//	}
//
//	headers := strings.Split(header, " ")
//	if len(headers) != 2 {
//		newErrorResponse(ctx, http.StatusBadRequest, "wrong header format")
//	}
//
//	token := headers[1]
//
//
//}
