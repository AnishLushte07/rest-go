package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"port"`
	DBName string `mapstructure:"dbName"`
}

var AppConfig *Config

func LoadConfig() {
	log.Println("...loading configuration")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	viper.Unmarshal(&AppConfig)

	if err != nil {
		log.Fatal(err)
	}
}
