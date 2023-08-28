package main

import (
	server "app/internal/server"
	flags "app/src/flags"
	"github.com/spf13/viper"
)

type BackendFlags struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Loglevel string            `yaml:"logLevel"`
	Backend  *BackendFlags     `yaml:"backend"`
	Redis    *flags.RedisFlags `yaml:"redis"`
}

func main() {
	v := viper.New()
	v.SetConfigName("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath("./src")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var conf Config
	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	server.SetupServer(conf.Redis).Run(conf.Backend.Host + ":" + conf.Backend.Port)
}
