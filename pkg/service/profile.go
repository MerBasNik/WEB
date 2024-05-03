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

func (s *ProfileService) CreateHobby(profId int, hobbies map[string][]chat.UserHobbyInput) ([]int, error) {
	return s.repo.CreateHobby(profId, hobbies)
}

func (s *ProfileService) GetAllHobby(profId int) ([]chat.UserHobby, error) {
	return s.repo.GetAllHobby(profId)
}

func (s *ProfileService) DeleteHobby(profId, hobbyId int) error {
	return s.repo.DeleteHobby(profId, hobbyId)
}

// func (s *ProfileService) GetAvatar(userId int, c *gin.Context) (string, error) {
// 	err = s.repo.Profile.GetAvatar(userId, id)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	if !user.ProfileIconPath.Valid {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
// 		return
// 	}

// 	avatarPath := settings.AppSettings.EnvVars.AvatarBasePath + "/" + user.ProfileIconPath.String
// 	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", user.ProfileIconPath.String))
// 	c.File(avatarPath)
// }



// func (s *ProfileService) RemoveAvatar(userId int, c *gin.Context) (string, error) {
// 	userId, err := getUserId(c)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	if user.ProfileIconPath.Valid {
// 		if err = os.Remove(settings.AppSettings.EnvVars.AvatarBasePath + "/" + user.ProfileIconPath.String); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove previous avatar"})
// 			return
// 		}
// 	}

// 	if err = user.RemoveProfileIconPath(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove avatar in DB"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Avatar removed successfully",
// 	})
// }
	







// func getEnv(key, fallback string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return fallback
// }

// func NewEnvVars() *EnvVars {
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}

// 	envVars := EnvVars{}

// 	envVars.AvatarBasePath = getEnv("AVATAR_BASE_PATH", "./")

// 	return &envVars
// }

// func Setup() {
// 	AppSettings.EnvVars = NewEnvVars()
// }