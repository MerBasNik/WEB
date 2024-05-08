package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type ChatItemService struct {
	repo     repository.ChatItem
	listRepo repository.ChatList
}

func NewChatItemService(repo repository.ChatItem, listRepo repository.ChatList) *ChatItemService {
	return &ChatItemService{repo: repo, listRepo: listRepo}
}

func (s *ChatItemService) CreateItem(userId, listId int, input chat.ChatItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.CreateItem(input)
}

// func (s *ChatItemService) GetAll(userId, listId int) ([]chat.ChatItem, error) {
// 	return s.repo.GetAll(userId, listId)
// }

// func (s *ChatItemService) GetById(userId, itemId int) (chat.ChatItem, error) {
// 	return s.repo.GetById(userId, itemId)
// }

// func (s *ChatItemService) Delete(userId, itemId int) error {
// 	return s.repo.Delete(userId, itemId)
// }

// func (s *ChatItemService) Update(userId, itemId int, input chat.UpdateItemInput) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}

// 	return s.repo.Update(userId, itemId, input)
// }
