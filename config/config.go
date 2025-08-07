package config

import (
	"fmt"
	
	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port int64
}

type Redis struct {
	Port     string
	Host     string
	Db       int
	Password string
}

type Db struct {
	Host string
	User string
	Pass string
	Name string
	Port int64
}


type App struct {
	Name    string
	Version int8
	Server  Server
	Db      Db
	Redis   Redis
}

func Config(path string) (*App, error) {
	viper.SetConfigName("config") // نام فایل بدون پسوند
	viper.AddConfigPath(path)     // مسیر فایل کانفیگ

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var appConfig App
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &appConfig, nil
}
