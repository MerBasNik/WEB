package handler

import (
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
)

// @Summary Find Users for chat
// @Security ApiKeyAuth
// @Tags find
// @Description find users for chat
// @ID find-user-by-time
// @Accept  json
// @Produce  json
// @Param input body chat.FindUserInput true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/find_chats_users [post]
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

	var mas_id []int
	mas_id = append(mas_id, userId)
	mas_id = append(mas_id, list_id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"finded user_id for chat": mas_id,
	})
}



// @Summary Find Users by hobby for chat
// @Security ApiKeyAuth
// @Tags find
// @Description find users by hobby for chat
// @ID find-user-by-hobby
// @Accept  json
// @Produce  json
// @Param input body chat.FindUserInput true "list info"
// @Success 200 {integer} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/find_chats_users_by_hobby [post]
func (h *Handler) findUsersByHobby(c *gin.Context)  {
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

	lists, err := h.services.ChatList.FindByHobby(userId, list_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllHobbyResponse{
		Data: lists,
	})
}

// @Summary Create Chat
// @Security ApiKeyAuth
// @Tags chats
// @Description create chat
// @ID create-chat
// @Accept  json
// @Produce  json
// @Param list_users_id body chat.UsersForChat true "Chat Users Id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/create_chat [post]
func (h *Handler) createList(c *gin.Context) {
	var list_users_id chat.UsersForChat
	if err := c.BindJSON(&list_users_id); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	chat_id, err := h.services.ChatList.Create(list_users_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"chat_id": chat_id,
	})
}


// @Summary Rename Chat
// @Security ApiKeyAuth
// @Tags chats
// @Description rename chat
// @ID rename-chat
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Param   chatName body chat.UpdateChat true "list info"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/rename_chat/{chat_id} [put]
func (h *Handler) renameChat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var chatName chat.UpdateChat
	if err := c.BindJSON(&chatName); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.ChatList.RenameChat(userId, chat_id, chatName)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
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

// @Summary Delete Find Users
// @Security ApiKeyAuth
// @Tags find
// @Description delete find users
// @ID delete-find-users
// @Accept  json
// @Produce  json
// @Param list_users_id body chat.UsersForChat true "Chat Users Id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/delete_find_users [delete]
func (h *Handler) deleteFindUsers(c *gin.Context) {
	var list_users_id chat.UsersForChat
	if err := c.BindJSON(&list_users_id); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ChatList.DeleteFindUsers(list_users_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
