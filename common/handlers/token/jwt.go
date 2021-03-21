package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go_api/common/handlers/config"
	"log"
	"time"
)

type Claims struct {
	UserName string
	PassWord string
	jwt.StandardClaims
}

/*
	@summary: 根据用户名和密码生成token
	@param: userName用户名，passWord密码
	@retur: token, err
*/
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zqb",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(config.Viper.GetString("jwt.secret"))
	token, err := tokenClaims.SignedString(secret)
	log.Println(token)
	return token, err

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Viper.GetString("jwt.secret")), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
