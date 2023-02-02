package mytumblrhandlers

import (
	"encoding/json"
	"errors"

	log "github.com/sirupsen/logrus"
)

func (t *MyTumblrHandler) GetBlogObj(blogName string) (interface{}, error) {
	if t == nil {
		log.Warn("handler not initialized")
	}
	if t.Client == nil {
		log.Fatal("client not initialized")
	}
	blogObj := t.Client.GetBlog(blogName)
	var err error
	if blogObj == nil {
		err = errors.New("blog doesn't exist")
	}
	return blogObj, err
}

// BlogRef is from t.GetBlogObj(blogName)
func (t *MyTumblrHandler) GetBlogInfo(blogName string) (string, error) {
	if t == nil {
		log.Warn("handler not initialized")
	}
	if t.Client == nil {
		log.Fatal("client not initialized")
	}
	blogRef := t.Client.GetBlog(blogName)
	blog, err := blogRef.GetInfo()
	if err != nil {
		log.Warn("could not retrieve blog reference")
		return "", err
	}
	blogInfo, err := json.MarshalIndent(blog, "", "\t")
	return string(blogInfo), err
}
