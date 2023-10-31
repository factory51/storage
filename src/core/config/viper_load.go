package config

import (
	"github.com/spf13/viper"
)

/*
	LoadConfig - Uses spf13 viper to load config file

[ACCEPTS]

	file_path string - path to folder containing config file
	file_name string - the name of the file without extention
	file_type string - type of config file to load json|ini|etc
*/
func LoadConfig(file_path string, file_name string, file_type string) (err error) {

	viper.AddConfigPath(file_path)
	viper.SetConfigName(file_name) // Register config file name (no extension)
	viper.SetConfigType(file_type) // Look for specific type
	err = viper.ReadInConfig()

	return

}
