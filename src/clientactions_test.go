package mytumblrhandlers

import (
	"path/filepath"
	"testing"
)

func TestGetBlogInfo(t *testing.T) {
	here, err := filepath.Abs(".")
	if err != nil {
		t.Fatal("could not find here: " + err.Error())
	}
	configPath := filepath.Join(here, "..", "cfg", "config.secret")
	ok := InitHandler(configPath)
	if !ok {
		t.Fatal("could not init handler")
	}
	_, err = GetBlogInfo("pukicho")
	if err != nil {
		t.Fatal("failed to get blog info: " + err.Error())
	}
}

func TestGetPostsThread(t *testing.T) {
	here, err := filepath.Abs(".")
	if err != nil {
		t.Fatal("could not find here: " + err.Error())
	}
	configPath := filepath.Join(here, "..", "cfg", "config.secret")
	ok := InitHandler(configPath)
	if !ok {
		t.Fatal("could not init handler")
	}
	blogObj, err := GetBlog("staff")
	if err != nil {
		t.Fatalf("failed to get blog: %s", err.Error())
	}
	posts, _, err := GetTextPostThread(blogObj, NOWTIME, DEFAULTLIMITNUMBER)
	if err != nil {
		t.Fatalf("failed to get latest post thread: %s", err.Error())
	}
	if len(posts) != DEFAULTLIMITNUMBER {
		t.Fatalf("amount of blog posts found is not the limit requested; \"staff\" blog should have more than 20 posts")
	}
}
