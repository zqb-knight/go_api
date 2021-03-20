package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Viper = viper.New()
)

func InitConfig(confPath string) {
	//默认配置文件
	if confPath == "" {
		Viper.SetConfigName("config")
		Viper.SetConfigType("toml")
		Viper.AddConfigPath("./conf")
	} else {
		Viper.SetConfigFile(confPath)
	}

	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
