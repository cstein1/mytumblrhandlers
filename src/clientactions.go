package mytumblrhandlers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tumblr/tumblr.go"
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

func (t *MyTumblrHandler) GetPosts(blogObj *tumblr.BlogRef) (string, error) {
	err := t.IsValid()
	if err != nil {
		return "", err
	}
	v := url.Values{}
	// TODO: these need lots of work... Offset required
	v.Add("before", time.Now().GoString())
	v.Add("limit", "3")
	log.Tracef("getting posts from blogref: %s", blogObj.Name)
	posts, err := blogObj.GetPosts(v)
	for _, post := range posts.Posts {
		fmt.Printf("POST: %s\n", post.Id)
		postObj := t.Client.GetPost(post.Id, blogObj.Name)
		if postObj != nil {
			fmt.Printf("post: %v\n", postObj)
		} else {
			fmt.Printf("No post\n")
		}
	}
	return "", nil
}
