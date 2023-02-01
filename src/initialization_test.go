package mytumblrhandlers

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestInitHandler(t *testing.T) {
	here, err := filepath.Abs(".")
	if err != nil {
		t.Fatal("could not find here")
	}
	configPath := filepath.Join(here, "..", "cfg", "config.secret")
	tumblrClient := *InitHandler(configPath)
	resp, err := tumblrClient.Client.GetDashboard()
	if err != nil {
		t.Fatal("could not get blog")
	}
	fmt.Printf("%v", resp)
}
