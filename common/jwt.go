package common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var screte string = "0115"

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(4 * time.Hour)

	claims := Claims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zqb",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(screte)

	return ss, err

}

func ParseToken(tokenString string) (*Claims, error) {
	toekn, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return screte, nil
	})
	if toekn != nil {
		if cliams, ok := toekn.Claims.(*Claims); ok && toekn.Valid {
			return cliams, nil
		}
	}

	return nil, err

}
