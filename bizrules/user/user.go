package bizrules

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"gorillatest/dal/mongodb/user"
	userdb "gorillatest/model/db"
	"gorillatest/model/frombackend"
	"gorillatest/model/tobackend"
	"gorillatest/utils"
)

func Login(userLogin tobackend.UserLogin) (*frombackend.User, error) {
	hashpassword, err := utils.HashPassword(userLogin.Password)
	if err != nil {
		return nil, err
	}

	userdb, err := user.GetByEmail(userLogin.Email)
	if err != nil {
		return nil, err
	}

	if userdb == nil {
		return nil, fmt.Errorf("Пользователь не найден")
	}

	if utils.CheckPasswordHash(hashpassword, userdb.HashPassword) {
		return nil, fmt.Errorf("Пароль или логин некорректен")
	}

	userFromBackend, err := fill(userdb)
	if err != nil {
		return nil, err
	}

	return userFromBackend, nil
}

func Register(userRegister tobackend.UserRegister) (*frombackend.User, error) {
	_, err := mail.ParseAddress(userRegister.Email)
	if err != nil {
		return nil, err
	}

	if len(userRegister.Password) < 3 {
		return nil, errors.New("Пароль меньше 3 символов")
	}

	if len(userRegister.UserName) == 0 {
		return nil, errors.New("Поле UserNamr не заполнено")
	}

	hashpassword, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return nil, err
	}

	var userdb userdb.User
	userdb.Email = userRegister.Email
	userdb.UserName = userRegister.UserName
	userdb.HashPassword = hashpassword

	err = user.Insert(&userdb)
	if err != nil {
		return nil, err
	}

	userFromBackend, err := fill(&userdb)
	if err != nil {
		return nil, err
	}

	return userFromBackend, nil
}

func GetUser(stoken string) (*frombackend.User, error) {
	if len(stoken) == 0 {
		return nil, fmt.Errorf("Не передан header authorization")
	}

	if !strings.HasPrefix(stoken, "Token ") {
		return nil, fmt.Errorf("Header authorization не начинается с Token ")
	}

	token := stoken[6:]
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("Некорректный токен ")
	}

	suserid, ok := (*claims)["userid"]
	if !ok {
		return nil, fmt.Errorf("Некорректный токен  Отсутствует userid")
	}

	userid, ok := suserid.(string)
	if !ok {
		return nil, fmt.Errorf("Некорректный токен  userid не строка")
	}

	userdb, err := user.GetById(userid)
	if err != nil {
		return nil, err
	}

	if userdb == nil {
		return nil, fmt.Errorf("Пользователь не найден")
	}

	userFromBackend, err := fill(userdb)
	if err != nil {
		return nil, err
	}

	return userFromBackend, nil
}

func fill(userdb *userdb.User) (*frombackend.User, error) {
	var userFromBackend frombackend.User

	userFromBackend.Id = userdb.Id
	userFromBackend.Email = userdb.Email
	userFromBackend.UserName = userdb.UserName

	if token, err := makeToken(userdb.Id); err != nil {
		return nil, err
	} else {
		userFromBackend.Token = token
	}

	return &userFromBackend, nil
}

func makeToken(userid string) (string, error) {
	dict := make(map[string]string)
	dict["userid"] = userid

	if token, err := utils.MakeToken(dict); err == nil {
		return token, nil
	} else {
		return "", err
	}
}
