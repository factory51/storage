package config

import (
	"github.com/spf13/viper"
)

func GetInt(key string) (val interface{}) {

	val = viper.GetInt(key)

	return
}
