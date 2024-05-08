package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type ChatListService struct {
	repo repository.ChatList
}

func NewChatListService(repo repository.ChatList) *ChatListService {
	return &ChatListService{repo: repo}
}

func (s *ChatListService) CreateList(requestCreateList chat.RequestCreateList) (int, error) {
	return s.repo.CreateList(requestCreateList)
}

func (s *ChatListService) RenameChat(userId, chatId int, chat chat.UpdateChat) error {
	return s.repo.RenameChat(userId, chatId, chat)
}

func (s *ChatListService) GetAllLists(userId int) ([]chat.ChatList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ChatListService) GetListById(userId, listId int) (chat.ChatList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *ChatListService) DeleteList(userId, listId int) error {
	return s.repo.DeleteList(userId, listId)
}

func (s *ChatListService) UpdateList(userId, listId int, input chat.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateList(userId, listId, input)
}

func (s *ChatListService) FindByTime(userId int, input chat.FindUserInput) ([]int, error) {
	return s.repo.FindByTime(userId, input)
}

func (s *ChatListService) FindThreeByHobby(list_users []int) ([]chat.UserHobby, error) {
	return s.repo.FindThreeByHobby(list_users)
}

func (s *ChatListService) FindTwoByHobby(list_users []int) ([]chat.UserHobby, error) {
	return s.repo.FindTwoByHobby(list_users)
}

func (s *ChatListService) DeleteFindUsers(userId chat.RequestCreateList) error {
	return s.repo.DeleteFindUsers(userId)
}