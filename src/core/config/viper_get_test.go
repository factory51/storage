package config

import (
	"testing" //testing suite
)

func TestGet(t *testing.T) {

	want := "Hello World"

	//load our config again for safety in testing
	err := LoadConfig("../../conf", "test.conf", "json")

	if err != nil {
		t.Errorf("Cannot read test config file cannot proceed.\nReason: %v\n", err.Error()) // bugger
	}

	got := Get("message")

	if got != want { //check
		t.Errorf("got %q, wanted %q", got, want) //bugger
	}
}
