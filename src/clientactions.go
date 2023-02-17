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

func (t *MyTumblrHandler) GetPosts(blogObj *tumblr.BlogRef, epoch, postType string, limit int) (postsOutput []*tumblr.Post, err error) {
	err = t.IsValid()
	if err != nil {
		return
	}
	limitStr := strconv.Itoa(limit)
	postsOutput = make([]*tumblr.Post, limit)
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
	postsInterface, err := postObj.All()
	if err != nil {
		log.Errorf("failed to interpret posts with error: %v", err.Error())
		return
	}
	for i, post := range postsInterface {
		postsOutput[i] = post.GetSelf()
	}
	return
}

func (t *MyTumblrHandler) GetPostsBody(blogObj *tumblr.BlogRef, epoch, postType string, limit int) (postsOutput []string, latestPostEpoch string, err error) {
	err = t.IsValid()
	if err != nil {
		return
	}
	postsOutput = make([]string, limit)
	postsInterface, err := t.GetPosts(blogObj, epoch, postType, limit)
	for i, post := range postsInterface {
		accessiblePost := post.GetSelf()
		postsOutput[i] = accessiblePost.Body
	}
	latestPostEpoch = postsInterface[limit-1].GetSelf().Date
	return
}

func (t *MyTumblrHandler) GetPostsThread(blogObj *tumblr.BlogRef, epoch, postType string, limit int) (postsOutput [][]string, latestPostEpoch string, err error) {
	err = t.IsValid()
	if err != nil {
		return
	}
	postsOutput = make([][]string, limit)
	postsInterface, err := t.GetPosts(blogObj, epoch, postType, limit)
	for i, post := range postsInterface {
		accessiblePost := post.GetSelf()
		postsTrail := accessiblePost.Trail
		textTrail := make([]string, len(postsTrail))
		for j, reply := range postsTrail {
			// reply.Content is subject to change: https://www.tumblr.com/docs/npf
			replyContent, err := ExtractTextHTML(reply.Content)
			if err != nil {
				log.Warnf("did not track text from a reply with err: %s", err.Error())
			}
			textTrail[j] = fmt.Sprintf("%s: %s", reply.Blog.Name, replyContent)
		}
		postsOutput[i] = textTrail
	}
	latestPostEpoch = postsInterface[limit-1].GetSelf().Date
	return
}
