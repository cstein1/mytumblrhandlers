package mytumblrhandlers

import (
	"path/filepath"
	"testing"
)

func TestInitHandler(t *testing.T) {
	here, err := filepath.Abs(".")
	if err != nil {
		t.Fatal("could not find here")
	}
	configPath := filepath.Join(here, "..", "cfg", "config.secret")
	ok := InitHandler(configPath)
	if !ok {
		t.Fatal("could not init handler")
	}
	resp, err := GetDashboard()
	if err != nil || resp == nil {
		t.Fatal("could not get blog")
	}
}
