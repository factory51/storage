package config

import (
	"testing" //testing suite
)

func TestLoadConfig(t *testing.T) {

	// test a known file
	err := LoadConfig("../../conf", "test.conf", "json")

	if err != nil {
		t.Errorf("Cannot read test config file cannot proceed.\nReason: %v\n", err.Error()) // bugger
	}

	//want := "Hello World"

}
