package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
)

//const maxAvatarSize = 10 * 1024 * 1024 // 10 MB

// @Summary Create Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description create profile
// @ID create-profile
// @Accept  json
// @Produce  json
// @Param input body chat.Profile true "profile info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/profile/create_profile [post]
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

// @Summary Edit Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description edit profile
// @ID edit-profile
// @Accept  json
// @Produce  json
// @Param   prof_id path int true "Prof Id"
// @Param   input body chat.UpdateProfile true "profile info"
// @Success 200 {integer} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/profile/edit_profile/{prof_id} [put]
func (h *Handler) editProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	prof_id, err := strconv.Atoi(c.Param("prof_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input chat.UpdateProfile
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Profile.EditProfile(userId, prof_id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

type getUserProfile struct {
	Data chat.Profile `json:"data"`
}

// @Summary Get Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description get profile
// @ID get-profile
// @Accept  json
// @Produce  json
// @Param   prof_id path int true "Prof Id"
// @Success 200 {integer} getUserProfile
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/profile/get_profile/{prof_id} [get]
func (h *Handler) getProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	prof_id, err := strconv.Atoi(c.Param("prof_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	profile, err := h.services.Profile.GetProfile(userId, prof_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserProfile{
		Data: profile,
	})
}

// @Summary      Upload Avatar
// @Security 	 ApiKeyAuth
// @Tags         profile
// @Description  updates user avatar
// @Id			 upload-avatar
// @Accept       multipart/form-data
// @Produce      json
// @Param 		 avatar formData file true "Avatar image file (JPG, JPEG, PNG, or GIF)"
// @Success      200  {object}  string
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/profile/upload_avatar [put]
func (h *Handler) uploadAvatar(c *gin.Context) {
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "no file uploaded")
		return
	}

	defer file.Close()

	var format = filepath.Ext(fileHeader.Filename)
	if !IsAvatarHasAllowedExtension(format) {
		NewErrorResponse(c, http.StatusBadRequest, "it is not JPG, JPEG, PNG, GIF")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// if fileHeader.Size <= maxAvatarSize && fileHeader.Size > 0 {
	// 	NewErrorResponse(c, http.StatusBadRequest, "avatar file size is too large")
	// 	return
	// }

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./avatars", os.ModePerm)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "error while mkdir ./avatars")
		return
	}

	// Create a new file in the uploads directory
	avatarFilename := fmt.Sprintf("./avatars/%d_%d_%s", time.Now().Unix(), userId, format)
	dst, err := os.Create(avatarFilename)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "error create new file")
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "did not copy files")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"directory": avatarFilename,
	})
}


// @Summary Create hobby
// @Security ApiKeyAuth
// @Tags hobby
// @Description create hobby
// @ID create-hobby
// @Accept  json
// @Produce  json
// @Param   prof_id path int true "Prof Id"
// @Param   input body map[string][]chat.UserHobbyInput true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/profile/{prof_id}/hobby/create_hobby [post]
func (h *Handler) createHobby(c *gin.Context) {
	profId, err := strconv.Atoi(c.Param("prof_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input map[string][]chat.UserHobbyInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hobby_id, err := h.services.Profile.CreateHobby(profId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"hobby_id": hobby_id,
	})
}

type getAllHobbyResponse struct {
	Data []chat.UserHobby `json:"data"`
}

// @Summary Get All Hobby
// @Security ApiKeyAuth
// @Tags hobby
// @Description get all hobby
// @ID get-all-hobby
// @Accept  json
// @Produce  json
// @Param   prof_id path int true "Prof Id"
// @Success 200 {object} getAllHobbyResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/profile/{prof_id}/hobby/get_hobby [get]
func (h *Handler) getAllHobby(c *gin.Context) {
	profId, err := strconv.Atoi(c.Param("prof_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	hobbylist, err := h.services.Profile.GetAllHobby(profId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllHobbyResponse{
		Data: hobbylist,
	})
}

// @Summary Delete Hobby
// @Security ApiKeyAuth
// @Tags hobby
// @Description delete hobby
// @ID delete-hobby
// @Accept  json
// @Produce  json
// @Param   prof_id path int true "Prof Id"
// @Param   hobby_id path int true "Hobby Id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/profile/{prof_id}/hobby/delete_hobby/{hobby_id} [delete]
func (h *Handler) deleteHobby(c *gin.Context) {
	hobby_id, err := strconv.Atoi(c.Param("hobby_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	profId, err := strconv.Atoi(c.Param("prof_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Profile.DeleteHobby(profId, hobby_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}


func IsAvatarHasAllowedExtension(ext string) bool {
	if ext == ".jpg" || ext == ".png" || ext == ".jpeg" || ext == ".gif" {
		return true
	}
	return false
}