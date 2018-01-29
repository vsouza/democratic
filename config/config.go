package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("../../")
	fmt.Println(v)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
}

func GetConfig() *viper.Viper {
	return config
}
