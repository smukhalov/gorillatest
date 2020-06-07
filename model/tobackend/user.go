package tobackend

type UserLoginFields struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	UserLoginFields `json:"user"`
}

type UserRegisterFields struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	UserRegisterFields `json:"user"`
}
