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
// @Param input body chat.FindUserInput true "Time"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/find_chats_users [post]
func (h *Handler) findUsersByTime(c *gin.Context) {
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

	list_id, users_info, err := h.services.ChatList.FindByTime(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ChatList.UpdateFindUsersTable(users_info, input.Count)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var lists []int
	if (input.Count == 2) {
		if (len(list_id) == 2) {
			lists = list_id
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"finded user_id for chat": 	lists,
		})
	}
	if (input.Count == 3) {
		if (len(list_id) == 3) {
			lists = list_id
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"finded user_id for chat": 	lists,
		})
	}
}

type getHobbyAndIdResponse struct {
	Data []chat.UserHobby `json:"data"`
	Users_id []int `json:"users_id"`
}

// @Summary Find Users by hobby for chat
// @Security ApiKeyAuth
// @Tags find
// @Description find users by hobby for chat
// @ID find-user-by-hobby
// @Accept  json
// @Produce  json
// @Param input body chat.FindUserInput true "Time"
// @Success 200 {integer} getHobbyAndIdResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/find_chats_users_by_hobby [post]
func (h *Handler) findUsersByHobby(c *gin.Context) {
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

	list_id, users_info, err := h.services.ChatList.FindByTime(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ChatList.UpdateFindUsersTable(users_info, input.Count)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var lists []chat.UserHobby
	var ids []int
	if (input.Count == 2) {
		if (len(list_id) == 2) {
			lists, err = h.services.ChatList.FindTwoByHobby(list_id)
			if err != nil {
				NewErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			ids = list_id
		}
		c.JSON(http.StatusOK, getHobbyAndIdResponse{
			Data: lists,
			Users_id: ids,
		})
	}
	if (input.Count == 3) {
		if (len(list_id) == 3) {
			lists, err = h.services.ChatList.FindThreeByHobby(list_id)
			if err != nil {
				NewErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			ids = list_id
		}
		c.JSON(http.StatusOK, getHobbyAndIdResponse{
			Data: lists,
			Users_id: ids,
		})
	}
}


// @Summary Create Chat
// @Security ApiKeyAuth
// @Tags chats
// @Description create chat
// @ID create-chat
// @Accept  json
// @Produce  json
// @Param input body chat.UsersForChat true "Chat Users Id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/create_chat [post]
func (h *Handler) createList(c *gin.Context) {
	var input chat.UsersForChat
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	chat_id, err := h.services.ChatList.CreateList(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ChatList.DeleteFindUsers(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"chat_id": chat_id,
	})
}
/*
type getAllListsResponse struct {
	Data []chat.ChatList `json:"data"`
}
*/

type getChatInfoResponse struct {
	Data []chat.UserHobby 	`json:"data"`
	UsersInfo map[int]string `json:"users_info"`
}

// @Summary Get Info For Chat By Id
// @Security ApiKeyAuth
// @Tags chats
// @Description get info chat by id
// @ID get-info-chat-by-id
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Success 200 {object} getChatInfoResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/get_info_for_chat/{chat_id} [get]
func (h *Handler) getInfoForListById(c *gin.Context) {
	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var users_id []int
	users_id, err = h.services.ChatList.GetUserByListId(chat_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var users_avatar []string
	users_avatar, err = h.services.ChatList.GetUserAvatar(users_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	users_info := make(map[int]string, len(users_id))
	for i := 0; i < len(users_id); i++ {
		users_info[users_id[i]] = users_avatar[i]
	}

	var lists []chat.UserHobby
	if (len(users_id) == 2) {
		lists, err = h.services.ChatList.FindTwoByHobby(users_id)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, getChatInfoResponse{
			Data: lists,
			UsersInfo: users_info,
		})
	}
	if (len(users_id) == 3) {
		lists, err = h.services.ChatList.FindThreeByHobby(users_id)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return			
		}
		c.JSON(http.StatusOK, getChatInfoResponse{
			Data: lists,
			UsersInfo: users_info,
		})
	}
}

// @Summary Rename Chat
// @Security ApiKeyAuth
// @Tags chats
// @Description rename chat
// @ID rename-chat
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Param input body chat.UpdateChat true "Chat Name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/rename_chat/{chat_id} [put]
func (h *Handler) renameChat(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input chat.UpdateChat
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.ChatList.RenameChat(user_id, chat_id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}


/*
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

	lists, err := h.services.ChatList.GetAllLists(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

type getListsResponse struct {
	Data chat.ChatList `json:"data"`
	Users_id []int `json:"users_id"`
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

	chat, err := h.services.ChatList.GetListById(userId, chat_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var users_id []int
	users_id, err = h.services.ChatList.GetUserByListId(chat_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getListsResponse{
		Data: chat,
		Users_id: users_id,
	})
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

	if err := h.services.ChatList.UpdateList(userId, chat_id, input); err != nil {
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

	err = h.services.ChatList.DeleteList(userId, chat_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
*/