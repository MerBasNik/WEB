package handler

import (
	"net/http"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body chat.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input chat.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// func (h *Handler) resetPassword(c *gin.Context) {
// 	var input chat.ResetPasswordInput

// 	if err := c.BindJSON(&input); err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
// 		return
// 	}

// 	err := h.services.Authorization.ResetPassword(input.Password)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, statusResponse{
// 		Status: "ok",
// 	})
// }


func (h *Handler) forgotPassword(c *gin.Context) {
	var inputEmail chat.ForgotPasswordInput

	if err := c.BindJSON(&inputEmail); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Authorization.ForgotPassword(inputEmail.Email)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input chat.ResetPasswordInput
	err = h.services.Authorization.ResetPassword(inputEmail.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// func (h *Handler) changePassword(c *gin.Context) {
// 	var input chat.ChangePasswordInput

// 	if err := c.BindJSON(&input); err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
// 		return
// 	}

// 	err := h.services.Authorization.PasswordChange(input)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, statusResponse{
// 		Status: "ok",
// 	})
// }