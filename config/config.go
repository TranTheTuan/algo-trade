package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfigKey(key string) (string, error) {
	key, ok := viper.Get(key).(string)
	if !ok {
		return "", errors.New("can not get config key")
	}
	return key, nil
}
