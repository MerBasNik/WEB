package chat

type User struct {
	Id       	 int    	`json:"-" db:"id"`
	Email	 	 string 	`json:"email" binding:"required"`
	Password 	 string 	`json:"password" binding:"required"`
}

type Profile struct {
	Name     	 string 	`json:"name" binding:"required"`
	Surname  	 string 	`json:"surname" binding:"required"`
	Photo 	 	 string 	`json:"photo" binding:"required"`
	Telegram 	 string 	`json:"telegram" binding:"required"`
	City 	 	 string 	`json:"city" binding:"required"`
	FindStatus 	 bool 		`json:"findstatus"`
}

type UpdateProfile struct {
	Name     	 *string 	`json:"name"`
	Surname  	 *string 	`json:"surname"`
	Photo 	 	 string 	`json:"photo"`
	Telegram 	 *string 	`json:"telegram"`
	City 	 	 *string 	`json:"city"`
}

type UsersHobbyList struct {
	Id int
	UserId int
	UserHobbyId int
}

type UserHobby struct {
	Id int 
	Description string `json:"description" binding:"required"`	
}

type ForgotPasswordInput struct {
	Email 	string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password 		string `json:"password" binding:"required"`
	PasswordRepeat  string `json:"password-repeat" binding:"required"`
}