package config

import (
	"github.com/spf13/viper"
)

func Get(key string) (val interface{}) {

	val = viper.Get(key)

	return
}
