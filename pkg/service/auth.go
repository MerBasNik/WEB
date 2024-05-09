package service

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
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

func (s *AuthService) GetUser(email, password string) (chat.User, error) {
	return s.repo.GetUser(email, generatePasswordHash(password))
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

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var number = []byte("0123456789")
var alphaNumeric = append(charset, number...)

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(b)
}

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *AuthService) ForgotPassword(input string) (string, error) {
	token := randStr(30)
	err := s.repo.SetUserToken(token, input)
	if err != nil {
		return "", err
	}

	PasswordResetToken := Encode(token)
	//PasswordResetAt := time.Now().Add(time.Minute * 15)
	err = SendEmail(input, PasswordResetToken)
	return PasswordResetToken, err
}

func (s *AuthService) ResetPassword(resetToken, password string) error {
	token, err := Decode(resetToken)
	if err != nil {
		return errors.New("token during decoding false")
	}

	user, err := s.repo.GetUserEmail(token)
	if err != nil {
		return errors.New("error for get user email")
	}

	err = s.repo.DeleteUserToken(user)
	if err != nil {
		return errors.New("error for delete user token")
	}

	passwordHash := generatePasswordHash(password)
	return s.repo.ResetPassword(user.Email, passwordHash)
}

func SendEmail(email, resetToken string) error {
	// sender data
	from := os.Getenv("FROM_EMAIL")   //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTP_PWD") // ex: "ieiemcjdkejspqz"
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
	body := "<body>Hello, you have received an email with information to change your password.<br>" +
		"To set a new password, you need to click on the link below.<br>" +
		"You will be redirected to the rndmCoffee company page,<br>" +
		"where you can successfully complete this operation.<br><br>" +
		"https://infotech12.eljur.ru/authorize" + resetToken +
		"<a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://infotech12.eljur.ru/authorize/" + resetToken +
		"\"><br><br>RESET PASSWORD</a><br><br> If you haven't forgotten your password, you can ignore this email.<br><br>" +
		"With all due respect to you,<br> the rndmCoffee team</body>"

	//fmt.Println("body:", body)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	//fmt.Println("message:", string(message))
	err := smtp.SendMail(address, auth, from, to, message)
	return err
}
