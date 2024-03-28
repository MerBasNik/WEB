package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"github.com/joho/godotenv"
	"log"


	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
	"github.com/gin-gonic/gin"
)

const maxAvatarSize = int64(2 * 1024 * 1024) // 2 MB

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

func (s *ProfileService) UploadAvatar(userId int, c *gin.Context) (string, error) {
	file, err := c.FormFile("avatar")
	if err != nil {
		return "", errors.New("no file uploaded")
	}

	if !IsAvatarInSizeRange(file.Size) {
		return "", errors.New("avatar file size is too large")
	}

	ext := filepath.Ext(file.Filename)
	if !IsAvatarHasAllowedExtension(ext) {
		return "", errors.New("invalid avatar file extension")
	}

	avatarFilename := fmt.Sprintf("avatar_%v_%d%s", time.Now().Unix(), userId, ext)
	avatarSavePath := AppSettings.EnvVars.AvatarBasePath + "/" + avatarFilename

	if err = os.Remove(AppSettings.EnvVars.AvatarBasePath + "/" + avatarFilename); err != nil {
		return "", errors.New("failed to delete previous avatar")
	}

	// if err = c.SaveUploadedFile(file, avatarSavePath); err != nil {
	// 	return "", errors.New("Failed to save file")
	// }

	return avatarSavePath, err
}




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
	


func IsAvatarInSizeRange(size int64) bool {
	return size <= maxAvatarSize && size > 0
}

func getAvatarAllowedExtensions() map[string]bool {
	return map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
}

func IsAvatarHasAllowedExtension(extension string) bool {
	if _, ok := getAvatarAllowedExtensions()[extension]; !ok {
		return false
	}

	return true
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewEnvVars() *EnvVars {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	envVars := EnvVars{}

	envVars.AvatarBasePath = getEnv("AVATAR_BASE_PATH", "./")

	return &envVars
}

func Setup() {
	AppSettings.EnvVars = NewEnvVars()
}