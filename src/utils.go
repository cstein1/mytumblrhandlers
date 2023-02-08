package mytumblrhandlers

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func init() {
	HandlerNotIntialized = errors.New("handler not initialized")
	ClientNotInitialized = errors.New("client not initialized")
}

var HandlerNotIntialized error
var ClientNotInitialized error

func (t *MyTumblrHandler) IsValid() (err error) {
	if t == nil {
		log.Warn("handler not initialized")
		err = HandlerNotIntialized
	}
	if t.Client == nil {
		log.Fatal("client not initialized")
		err = ClientNotInitialized
	}
	return err
}