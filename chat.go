package chat

import "errors"

type ChatList struct {
	UserId      int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ChatItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type FindUserInput struct {
	StartDay 	string `json:"startday" db:"startday"`
	EndDay 		string `json:"endday" db:"endday"`
	StartTime 	string `json:"starttime" db:"starttime"`
	EndTime 	string `json:"endtime" db:"endtime"`
}

type ItemLists struct {
	Id     int
	ChatListId int
	ChatItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}