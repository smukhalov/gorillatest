package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret []byte

func init() {
	hmacSampleSecret = []byte("Самый большой секрет!")
	// Load sample key data
	// if keyData, e := ioutil.ReadFile("test/hmacTestKey"); e == nil {
	// 	hmacSampleSecret = keyData
	// } else {
	// 	panic(e)
	// }
}

func MakeToken(dict map[string]string) (string, error) {
	if _, ok := dict["userid"]; !ok {
		return "", fmt.Errorf("Отсутствет ключ userid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": dict["userid"],
		"nbf":    time.Date(2015, 1, 1, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":    time.Date(2030, 1, 1, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// var hmacSampleSecret []byte
	// // Sign and get the complete encoded token as a string using the secret
	return token.SignedString(hmacSampleSecret)
}

func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
		//fmt.Println("Valid", claims["foo"], claims["exp"], claims["nbf"], claims["jti"])
	}
	return nil, fmt.Errorf("Token не валидный")
}
