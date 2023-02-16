package mytumblrhandlers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

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

func (t *MyTumblrHandler) GetPosts(blogObj *tumblr.BlogRef, epoch, postType string, limit int) (postsOutput []string, latestPostEpoch string, err error) {
	err = t.IsValid()
	if err != nil {
		return
	}
	limitStr := strconv.Itoa(limit)
	postsOutput = make([]string, limit)
	v := url.Values{}
	v.Add("before", epoch)
	v.Add("limit", limitStr)
	v.Add("type", postType)
	log.Tracef("getting posts from blogref: %s", blogObj.Name)
	postObj, err := blogObj.GetPosts(v)
	if err != nil {
		log.Errorf("failed to retrieve posts with error: %v", err.Error())
		return
	}
	postResponses, err := postObj.All()
	if err != nil {
		log.Errorf("failed to interpret posts with error: %v", err.Error())
		return
	}
	for i, post := range postResponses {
		accessiblePost := post.GetSelf()
		fmt.Printf("%v\n", accessiblePost.Summary)
		postsOutput[i] = accessiblePost.Summary
	}
	latestPostEpoch = postResponses[limit-1].GetSelf().Date
	return
}
