package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"time"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "asuSDFvsrjr23bjbb2r28qsq"
	signingKey = "834hucbqp*&#)bprrqcibSGu#Rprn;"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user chat.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ForgotPassword(input string) error {
	err := SendEmail(input)
	return err
}

func (s *AuthService) ResetPassword(email, password string) error {
	passwordHash := generatePasswordHash(password)
	return s.repo.ResetPassword(email, passwordHash)
}

func SendEmail(email string) error {
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address privided through toEmail argument
	to := []string{email}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Email Verification\n"
	// localhost:8080 will be removed by many email service but works with online sites
	// https must be used since we are sending personal data through url parameters

	//  ???????
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"http://localhost:3000/main\">Reset Password</a></body>"

	fmt.Println("body:", body)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	fmt.Println("message:", string(message))
	err := smtp.SendMail(address, auth, from, to, message)
	return err
}
