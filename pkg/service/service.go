package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	ForgotPassword(input string) (string, error)
	ResetPassword(resetToken, password string) error 
}

type Profile interface {
	CreateProfile(userId int, profile chat.Profile) (int, error)
	GetProfile(userId, profileId int) (chat.Profile, error)
	EditProfile(userId, profileId int, input chat.UpdateProfile) error

	CreateHobby(profId int, hobby chat.UserHobby) (int, error)
	GetAllHobby(profId int) ([]chat.UserHobby, error)
	DeleteHobby(profId, hobbyId int) error
	//UploadAvatar(profileId int, directory string) error
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
	Profile
	ChatList
	ChatItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Profile: 	   NewProfileService(repos.Profile),
		ChatList:      NewChatListService(repos.ChatList),
		ChatItem:      NewChatItemService(repos.ChatItem, repos.ChatList),
	}
}