package token

import (
	"github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"time"
	"errors"
)

const (
    SecretKey = "MBControlGroup dada"
)

func Generate(admin_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": admin_id,
		"exp": (time.Now().Add(time.Minute * 60 * 24 * 2)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	
	fmt.Println(token.Valid)
	return tokenString, err
}

func Valid(tokenString string) (string, error) {
	token1, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if token1 == nil {
		return "", errors.New("InValid token!")
	}
	if claims, ok := token1.Claims.(jwt.MapClaims); ok && token1.Valid {
		fmt.Println(3)
		return fmt.Sprintf("%v", claims["id"]), nil
	} else {
		return "", errors.New("InValid token!")
	}

}
