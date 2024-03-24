package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(userId int, profile chat.Profile) (int, error) {
	return s.repo.CreateProfile(userId, profile)
}

func (s *ProfileService) EditProfile(userId, profileId int, input chat.UpdateProfile) error {
	return s.repo.EditProfile(userId, profileId, input)
}

func (s *ProfileService) GetProfile(userId, profileId int) (chat.Profile, error) {
	return s.repo.GetProfile(userId, profileId)
}

func (s *ProfileService) CreateHobby(userId int, hobby chat.UserHobby) (int, error) {
	return s.repo.CreateHobby(userId, hobby)
}

func (s *ProfileService) GetAllHobby(userId int) ([]chat.UserHobby, error) {
	return s.repo.GetAllHobby(userId)
}

func (s *ProfileService) DeleteHobby(userId, hobbyId int) error {
	return s.repo.DeleteHobby(userId, hobbyId)
}

	
	