package main

import (
	server "app/internal/server"
	"fmt"
	"github.com/spf13/viper"
	// "log"
)

type RedisParams struct {
	Port     int
	User     string
	Password string
}

type BackendParams struct {
	Adress string
	Port   string
}

type Config struct {
	Loglevel string        `yaml:"logLevel"`
	Redis    RedisParams   `yaml:"redis"`
	Backend  BackendParams `yaml:"backend"`
}

func main() {
	v := viper.New()
	v.SetConfigName("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var conf Config

	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello world!")

	fmt.Println(conf.Loglevel)
	fmt.Println(conf.Redis.User)
	fmt.Println(conf.Redis.Port)
	fmt.Println(conf.Redis.Password)
	fmt.Println(conf.Redis)

	port := conf.Backend.Port
	adress := conf.Backend.Adress

	fmt.Println(adress, " ", port)

	server.SetupServer().Run(adress + ":" + port)
}
