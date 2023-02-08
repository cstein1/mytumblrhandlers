package mytumblrhandlers

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func (t *MyTumblrHandler) GetBlogObj(blogName string) (interface{}, error) {
	err := t.IsValid()
	if err != nil {
		return nil, err
	}
	blogObj := t.Client.GetBlog(blogName)
	if blogObj == nil {
		err = BlogDoesntExist
	}
	return blogObj, err
}

// BlogRef is from t.GetBlogObj(blogName)
func (t *MyTumblrHandler) GetBlogInfo(blogName string) (string, error) {
	err := t.IsValid()
	if err != nil {
		return "", err
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
