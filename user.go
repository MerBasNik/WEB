package chat

import "errors"

type UserInput struct {
	Email	 	 string 	`json:"email" db:"email" binding:"required"`
	Password 	 string 	`json:"password" db:"password" binding:"required"`
}

type User struct {
	Id       	 int    	`json:"id" db:"id"`
	Email	 	 string 	`json:"email" db:"email" binding:"required"`
	Password 	 string 	`json:"password" db:"password" binding:"required"`
	Token		 string     `json:"token" db:"token"`
}

type Profile struct {
	Id       	 int    	`json:"id" db:"id"`
	Name     	 string 	`json:"name" db:"name" binding:"required"`
	Surname  	 string 	`json:"surname" db:"surname" binding:"required"`
	Email  	 	 *string 	`json:"email" db:"email"`
	Photo 	 	 string 	`json:"photo" db:"photo" binding:"required"`
	Telegram 	 string 	`json:"telegram" db:"telegram" binding:"required"`
	Country		 string		`json:"country" db:"country"`
	City 	 	 string 	`json:"city" db:"city"`
	Birthday 	 string 	`json:"birthday" db:"birthday"`
}

type UpdateProfile struct {
	Name     	 *string 	`json:"name"`
	Surname  	 *string 	`json:"surname"`
	Email  		 *string 	`json:"email"`
	Photo 	 	 *string 	`json:"photo"`
	Telegram 	 *string 	`json:"telegram"`
	Country		 *string	`json:"country"`
	City 	 	 *string 	`json:"city"`
	Birthday 	 *string 	`json:"birthday"`
}

func (i UpdateProfile) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Email == nil && i.Photo == nil && i.Telegram == nil && i.City == nil && i.Birthday == nil && i.Country == nil{
		return errors.New("update structure has no values")
	}

	return nil
}


type UsersHobbyList struct {
	Id 			int
	UserId 		int
	UserHobbyId int
}

type UserHobby struct {
	Id 			int    `json:"id" db:"id"`
	Description string `json:"description" db:"description" binding:"required"`	
}

type UserHobbyInput struct {
	Description string `json:"description" db:"description" binding:"required"`	
}

type ForgotPasswordInput struct {
	Email 	string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password 		string `json:"password" binding:"required"`
	PasswordRepeat  string `json:"password-repeat" binding:"required"`
}