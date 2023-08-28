package main

import (
	server "app/internal/server"
	flags "app/src/flags"
	"github.com/spf13/viper"
)

type Config struct {
	Loglevel string            `yaml:"logLevel"`
	Redis    *flags.RedisFlags `yaml:"redis"`
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

	server.SetupServer(conf.Redis).Run()
}
