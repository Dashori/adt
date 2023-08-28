package main

import (
	server "app/internal/server"
	flags "app/src/flags"
	"fmt"
	"github.com/spf13/viper"
	"context"
	"time"
	// "log"
)

type Config struct {
	Loglevel string       `yaml:"logLevel"`
	Redis    *flags.RedisFlags   `yaml:"redis"`
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

	fmt.Println(conf.Loglevel)

	fmt.Println(conf.Redis)


	client, err := flags.NewRedisClient(conf.Redis)

	if err != nil {
		panic(err)
	}

	err = client.Set(context.TODO(), "key", "dasha", 1*time.Minute)

	if err != nil {
		panic(err)
	}

	ans, err := client.Get(context.TODO(), "key")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(ans[:]))


	server.SetupServer().Run()

}
