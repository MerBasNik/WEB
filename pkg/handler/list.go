package handler

import (
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
)

func (h *Handler) findUsersByTime(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input chat.FindUserInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list_id, err := h.services.ChatList.FindByTime(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"finded user_id for chat": list_id,
	})
}


// @Summary Create chat
// @Security ApiKeyAuth
// @Tags chats
// @Description create chat
// @ID create-chat
// @Accept  json
// @Produce  json
// @Param input body chat.ChatList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/create_chat [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input chat.ChatList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	chat_id, err := h.services.ChatList.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"chat_id": chat_id,
	})
}

type getAllListsResponse struct {
	Data []chat.ChatList `json:"data"`
}

// @Summary Get All Chats
// @Security ApiKeyAuth
// @Tags chats
// @Description get all chats
// @ID get-all-chats
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/get_all_chats [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.ChatList.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get Chat By Id
// @Security ApiKeyAuth
// @Tags chats
// @Description get chat by id
// @ID get-chat-by-id
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Success 200 {object} chat.ChatList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/get_chat/{chat_id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.ChatList.GetById(userId, chat_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}


// @Summary Update Chat
// @Security ApiKeyAuth
// @Tags chats
// @Description update chat
// @ID update-chat
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Param input body chat.UpdateListInput true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/update_chat/{chat_id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input chat.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.ChatList.Update(userId, chat_id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Chat
// @Security ApiKeyAuth
// @Tags chats
// @Description delete chat
// @ID delete-chat
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/delete_chat/{chat_id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.ChatList.Delete(userId, chat_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
