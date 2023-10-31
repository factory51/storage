package config

import (
	"path"
	"strings"
)

/*
ParseViperConfigPath - splits the supplied file path into the components used by viper to load a config file. NOTE assumes extention is the file type for ease of use
[ACCEPTS]

	input_path string - path to the file we wish to load with viper

[RETURNS]

	file_path string - the path to the directory the config file exists in
	file_name string - the name of the file without extention
	file_type string - the type of file we're going to load. Assumption extention matches type
*/
func ParseViperConfigPath(input_path string) (file_path string, file_name string, file_type string) {

	file_path = path.Dir(input_path)
	base := path.Base(input_path)
	file_name = base[:len(base)-len(path.Ext(base))]
	file_type = strings.Replace(path.Ext(input_path), ".", "", 1)

	return
}
