package chat

type RequestCreateList struct {
	Title   string `json:"title"`
	UsersId []int  `json:"usersId"`
}
