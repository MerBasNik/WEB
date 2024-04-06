package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"token": token,
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


// @Summary Reset Password
// @Tags auth
// @Description reset password
// @ID reset-password
// @Accept  json
// @Produce  json
// @Param   token path string true "Token"
// @Param   input body chat.ResetPasswordInput true "passwords"
// @Success 200 {string} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/reset-password/{token} [put]
func (h *Handler) resetPassword(c *gin.Context) {
	token, err := strconv.Atoi(c.Param("token"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input chat.ResetPasswordInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, err := h.services.Authorization.ParseToken(fmt.Sprint(token))
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.Authorization.ResetPassword(userId, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Forgot Password
// @Tags auth
// @Description forgot password
// @ID forgot-password
// @Accept  json
// @Produce  json
// @Param input body chat.ForgotPasswordInput true "email"
// @Success 200 {string} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/forgot-password [post]
func (h *Handler) forgotPassword(c *gin.Context) {
	var input chat.ForgotPasswordInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.services.Authorization.GetUserId(input.Email)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.Authorization.ForgotPassword(input.Email, user.Id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
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