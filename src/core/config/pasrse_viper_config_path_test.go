package config

import (
	"testing"
)

func TestParseViperConfigPath(t *testing.T) {

	test_path := "./var/vhosts/www/data/app.config.json"

	file_path, file_name, file_type := ParseViperConfigPath(test_path)

	if file_path != "var/vhosts/www/data" {
		t.Errorf("Cannot get file path\n") // bugger
	}

	if file_name != "app.config" {
		t.Errorf("Cannot get file name without suffix\n") // bugger
	}

	if file_type != "json" {
		t.Errorf("Cannot get file type\n") // bum
	}

	//test default working path

	test_path = "./conf/app.config.json"

	file_path, file_name, file_type = ParseViperConfigPath(test_path)

	if file_path != "conf" {
		t.Errorf("Cannot get file path: got:\n") // bugger
	}

	if file_name != "app.config" {
		t.Errorf("Cannot get file name without suffix\n") // bugger
	}

	if file_type != "json" {
		t.Errorf("Cannot get file type\n") // bum
	}
}
