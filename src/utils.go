package mytumblrhandlers

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	HandlerNotIntialized = errors.New("handler not initialized")
	ClientNotInitialized = errors.New("client not initialized")
	BlogDoesntExist = errors.New("blog doesn't exist")

	NOWTIME = fmt.Sprintf("%d", time.Now().Unix())
	DEFAULTPOSTTYPE = "text"
	DEFAULTLIMITNUMBER = 3
}

var NOWTIME string
var DEFAULTPOSTTYPE string
var DEFAULTLIMITNUMBER int

var HandlerNotIntialized error
var ClientNotInitialized error
var BlogDoesntExist error

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
