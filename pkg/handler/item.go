package handler

import (
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
	
)


// @Summary Create Chat Item
// @Security ApiKeyAuth
// @Tags chats items
// @Description create chat item
// @ID create-chat-item
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Param input body chat.ChatItem true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/{chat_id}/items/create_item [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat_id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input chat.ChatItem
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ChatItem.CreateItem(userId, chat_id, input.Username, input.Description, input.Chatlist_id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// type getAllListsItemsResponse struct {
// 	Data []chat.ChatItem `json:"data"`
// }

// @Summary Get All Chats Items
// @Security ApiKeyAuth
// @Tags chats items
// @Description get all chats items
// @ID get-all-chats-items
// @Accept  json
// @Produce  json
// @Param   chat_id path int true "Chat Id"
// @Success 200 {object} getAllListsItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/chats/{chat_id}/items/get_all_items [get]
// func (h *Handler) getAllItems(c *gin.Context) {
// 	userId, err := getUserId(c)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	chat_id, err := strconv.Atoi(c.Param("chat_id"))
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
// 		return
// 	}

// 	items, err := h.services.ChatItem.GetAll(userId, chat_id)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, getAllListsItemsResponse{
// 		Data: items,
// 	})
// }

// @Summary Get Chat Item By Id
// @Security ApiKeyAuth
// @Tags chats items
// @Description get chat item by id
// @ID get-chat-item-by-id
// @Accept  json
// @Produce  json
// @Param   item_id path int true "Item Id"
// @Success 200 {object} chat.ChatItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/get_item/{item_id} [get]
// func (h *Handler) getItemById(c *gin.Context) {
// 	userId, err := getUserId(c)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	item_id, err := strconv.Atoi(c.Param("item_id"))
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
// 		return
// 	}

// 	item, err := h.services.ChatItem.GetById(userId, item_id)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, item)
// }

// @Summary Update Chat Item
// @Security ApiKeyAuth
// @Tags chats items
// @Description update chat item
// @ID update-chat-item
// @Accept  json
// @Produce  json
// @Param   item_id path int true "Item Id"
// @Param input body chat.UpdateItemInput true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/update_item/{item_id} [put]
// func (h *Handler) updateItem(c *gin.Context) {
// 	userId, err := getUserId(c)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	item_id, err := strconv.Atoi(c.Param("item_id"))
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	var input chat.UpdateItemInput
// 	if err := c.BindJSON(&input); err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	if err := h.services.ChatItem.Update(userId, item_id, input); err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, statusResponse{"ok"})
// }

// @Summary Delete Chat Item
// @Security ApiKeyAuth
// @Tags chats items
// @Description delete chat item
// @ID delete-chat-item
// @Accept  json
// @Produce  json
// @Param   item_id path int true "Item Id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/delete_item/{item_id} [delete]
// func (h *Handler) deleteItem(c *gin.Context) {
// 	userId, err := getUserId(c)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	item_id, err := strconv.Atoi(c.Param("item_id"))
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
// 		return
// 	}

// 	err = h.services.ChatItem.Delete(userId, item_id)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, statusResponse{"ok"})
// }
