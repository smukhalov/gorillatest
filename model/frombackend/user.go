package frombackend

type UserFields struct {
	Id       string `json:"id"`
	UserName string `json:"usertname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type User struct {
	UserFields `json:"user"`
}
