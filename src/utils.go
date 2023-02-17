package mytumblrhandlers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func init() {
	ClientNotInitialized = errors.New("client not initialized")
	BlogDoesntExist = errors.New("blog doesn't exist")

	NOWTIME = fmt.Sprintf("%d", time.Now().Unix())
	DEFAULTPOSTTYPE = "text"
	DEFAULTLIMITNUMBER = 3
	DEFAULTOFFSET = 0
	// See https://www.tumblr.com/docs/en/api/v2#posts--retrieve-published-posts
	MAXIMUMLIMIT = 20
}

var NOWTIME string
var DEFAULTPOSTTYPE string
var DEFAULTLIMITNUMBER int
var MAXIMUMLIMIT int
var DEFAULTOFFSET int

var ClientNotInitialized error
var BlogDoesntExist error

func IsValid() (err error) {
	if client == nil {
		log.Fatal("client not initialized")
		err = ClientNotInitialized
	}
	return err
}

func ExtractTextHTML(inp string) (out string, err error) {
	body := "<mytumblrhandlers>" + inp + "</mytumblrhandlers>"
	reader := strings.NewReader(body)
	z := html.NewTokenizer(reader)
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err = z.Err()
			return
		case html.TextToken:
			out += string(z.Text()) + "\n"
		case html.StartTagToken, html.EndTagToken:
			if imgText := _helperGetImageTextFromHTML(z); len(imgText) > 0 {
				out += fmt.Sprintf(" %s ", imgText)
			}
			if tt == html.StartTagToken {
				depth++
			} else {
				depth--
				if depth == 0 {
					return
				}
			}
		case html.SelfClosingTagToken:
			// if an image is present, we want to link it in the return
			out += fmt.Sprintf(" %s ", _helperGetImageTextFromHTML(z))
		}
	}
}

func _helperGetImageTextFromHTML(z *html.Tokenizer) string {
	tagName, hasAttr := z.TagName()
	if string(tagName) == "img" && hasAttr {
		more := true
		var key, val []byte
		for more {
			key, val, more = z.TagAttr()
			if string(key) == "src" {
				return string(val)
			}
		}
	}
	return ""
}
