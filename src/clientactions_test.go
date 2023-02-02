package mytumblrhandlers

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGetBlogInfo(t *testing.T) {
	here, err := filepath.Abs(".")
	if err != nil {
		t.Fatal("could not find here: " + err.Error())
	}
	configPath := filepath.Join(here, "..", "cfg", "config.secret")
	tumblrClient := *InitHandler(configPath)
	info, err := tumblrClient.GetBlogInfo("pukicho")
	if err != nil {
		t.Fatal("failed to get blog info: " + err.Error())
	}
	fmt.Printf("%v", info)
}
