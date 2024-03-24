package handler

import (
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
)

// type ProfileInput struct {
// 	Name     	 string 	`json:"name" binding:"required"`
// 	Surname  	 string 	`json:"surname" binding:"required"`
// 	//Photo 	 	 string 	`json:"photo" binding:"required"`
// 	Telegram 	 string 	`json:"telegram" binding:"required"`
// 	City 	 	 string 	`json:"city" binding:"required"`
// 	//Hobby 	 	 []string 	`json:"hobby" binding:"required"`
// }

func (h *Handler) createProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input chat.Profile
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	profile_id, err := h.services.Profile.CreateProfile(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"profile_id": profile_id,
	})
}


func (h *Handler) editProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input chat.UpdateProfile
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Profile.EditProfile(userId, id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}