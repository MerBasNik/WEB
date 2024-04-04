package repository

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GetUser(email, password string) (chat.User, error)
	ResetPassword(email, password string) error 
}

type Profile interface {
	CreateProfile(userId int, profile chat.Profile) (int, error)
	GetProfile(userId, profileId int) (chat.Profile, error)
	EditProfile(userId, profileId int, input chat.UpdateProfile) error
	
	CreateHobby(userId int, hobby chat.UserHobby) (int, error)
	GetAllHobby(userId int) ([]chat.UserHobby, error)
	DeleteHobby(userId, hobbyId int) error
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
	Create(listId int, item chat.ChatItem) (int, error)
	GetAll(userId, listId int) ([]chat.ChatItem, error)
	GetById(userId, itemId int) (chat.ChatItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input chat.UpdateItemInput) error
}

type Repository struct {
	Authorization
	Profile
	ChatList
	ChatItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Profile: 	   NewProfilePostgres(db),
		ChatList:      NewChatListPostgres(db),
		ChatItem:      NewChatItemPostgres(db),
	}
}