package toekn

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
func GenerateToken(userName, passWord string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		userName,
		passWord,
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
	log.Println("未开发")
	return nil, nil
}
