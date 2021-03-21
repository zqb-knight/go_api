package model

import "log"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{UserName: username, PassWord: password}).First(&auth)
	if auth.ID > 0 {

		return true
	}
	log.Println("database query auth error")
	return false
}
