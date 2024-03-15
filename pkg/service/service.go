package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ChatList interface {
	Create(userId int, list chat.ChatList) (int, error)
	GetAll(userId int) ([]chat.ChatList, error)
	GetById(userId, listId int) (chat.ChatList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input chat.UpdateListInput) error
}

type ChatItem interface {
	Create(userId, listId int, item chat.ChatItem) (int, error)
	GetAll(userId, listId int) ([]chat.ChatItem, error)
	GetById(userId, itemId int) (chat.ChatItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input chat.UpdateItemInput) error
}

type Service struct {
	Authorization
	ChatList
	ChatItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ChatList:      NewChatListService(repos.ChatList),
		ChatItem:      NewChatItemService(repos.ChatItem, repos.ChatList),
	}
}