package chat

import "errors"

type User struct {
	Id       	 int    	`json:"id" db:"id"`
	Email	 	 string 	`json:"email" db:"email" binding:"required"`
	Password 	 string 	`json:"password" db:"password" binding:"required"`
}

type Profile struct {
	Id       	 int    	`json:"id" db:"id"`
	Name     	 string 	`json:"name" db:"name" binding:"required"`
	Surname  	 string 	`json:"surname" db:"surname" binding:"required"`
	Photo 	 	 string 	`json:"photo" db:"photo" binding:"required"`
	Telegram 	 string 	`json:"telegram" db:"telegram" binding:"required"`
	City 	 	 string 	`json:"city" db:"city" binding:"required"`
	FindStatus 	 bool 		`json:"findstatus" db:"findstatus"`
}

type UpdateProfile struct {
	Name     	 *string 	`json:"name"`
	Surname  	 *string 	`json:"surname"`
	Photo 	 	 *string 	`json:"photo"`
	Telegram 	 *string 	`json:"telegram"`
	City 	 	 *string 	`json:"city"`
}

func (i UpdateProfile) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Photo == nil && i.Telegram == nil && i.City == nil{
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

type ForgotPasswordInput struct {
	Email 	string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password 		string `json:"password" binding:"required"`
	PasswordRepeat  string `json:"password-repeat" binding:"required"`
}