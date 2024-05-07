package repository

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GetUser(email, password string) (chat.User, error)
	GetUserEmail(token string) (chat.User, error)
	ResetPassword(email, password string) error
	DeleteUserToken(user chat.User) error
	SetUserToken(token, email string) error 
}

type Profile interface {
	CreateProfile(userId int, profile chat.Profile) (int, error)
	GetProfile(userId, profileId int) (chat.Profile, error)
	EditProfile(userId, profileId int, input chat.UpdateProfile) error
	InitAllHobbies() error
	CreateHobby(profId int, hobbies map[string][]chat.UserHobbyInput) ([]int, error)
	GetAllHobby(profId int) ([]chat.UserHobby, error)
	DeleteHobby(profId, hobbyId int) error
	//UploadAvatar(profileId int, directory string) error
}

type ChatList interface {
	Create(userId chat.UsersForChat) (int, error)
	RenameChat(userId, chatId int, chat chat.UpdateChat) error
	GetAll(userId int) ([]chat.ChatList, error)
	GetById(userId, listId int) (chat.ChatList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input chat.UpdateListInput) error
	FindByTime(userId int, input chat.FindUserInput) ([]int, error)
	FindThreeByHobby(list_users []int) ([]chat.UserHobby, error)
	FindTwoByHobby(list_users []int) ([]chat.UserHobby, error)
	DeleteFindUsers(userId chat.UsersForChat) error
}

type ChatItem interface {
	GetUsers(userId, listId int) ([]int, error)
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