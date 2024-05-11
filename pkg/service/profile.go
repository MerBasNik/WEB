package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}
type EnvVars struct {
	AvatarBasePath string
}
type Settings struct {
	EnvVars *EnvVars
}

var AppSettings = &Settings{}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(userId int, profile chat.Profile) (int, error) {
	return s.repo.CreateProfile(userId, profile)
}

func (s *ProfileService) EditProfile(userId, profileId int, input chat.UpdateProfile) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.EditProfile(userId, profileId, input)
}

func (s *ProfileService) GetProfile(userId, profileId int) (chat.Profile, error) {
	return s.repo.GetProfile(userId, profileId)
}

func (s *ProfileService) InitAllHobbies() error {
	return s.repo.InitAllHobbies()
}

func (s *ProfileService) CreateHobby(profId int, hobbies []chat.UserHobbyInput) ([]int, error) {
	return s.repo.CreateHobby(profId, hobbies)
}

func (s *ProfileService) GetAllHobby(profId int) ([]chat.UserHobby, error) {
	return s.repo.GetAllHobby(profId)
}

func (s *ProfileService) DeleteHobby(profId, hobbyId int) error {
	return s.repo.DeleteHobby(profId, hobbyId)
}

func (s *ProfileService) GetProfileId(userId int) (int, error) {
	return s.repo.GetProfileId(userId)
} 