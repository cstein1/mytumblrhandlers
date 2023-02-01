package main

import (
	mytumblrhandlers "github.com/cstein1/mytumblrhandlers/src"
)

var FIRSTRUN = true
var DEFAULTCONFIGLOCATION = "./src/config.secret"

func main() {
	if FIRSTRUN {
		mytumblrhandlers.GetAccessToken(DEFAULTCONFIGLOCATION)
	} else {
		mytumblrhandlers.GetOAuthToken(DEFAULTCONFIGLOCATION)
	}
}
