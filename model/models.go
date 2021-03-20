package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_api/common/handlers/config"
	"log"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

var (
	db *gorm.DB
)

func Init() {
	var err error
	dbType := config.Viper.GetString("database.type")
	host := config.Viper.GetString("database.host")
	user := config.Viper.GetString("database.user")
	password := config.Viper.GetString("database.password")
	dbName := config.Viper.GetString("database.dbName")
	tablePrefix := config.Viper.GetString("database.tablePrefix")
	test := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	log.Println(test)
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println("数据库打开失败", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	if db != nil {
		defer db.Close()
	}
}
