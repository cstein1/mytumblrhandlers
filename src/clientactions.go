package mytumblrhandlers

import (
	log "github.com/sirupsen/logrus"
)

func (t *MyTumblrHandler) GetBlog(urlBase string) (interface{}, error) {
	if t == nil {
		log.Warn("handler not initialized")
	}
	resp, err := t.Client.Get("https://api.tumblr.com/v2/blog/" + urlBase)
	if err != nil {
		log.Warn("could not return blog with error: " + err.Error())
		return nil, err
	}
	return resp, err
}
