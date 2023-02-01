package mytumblrhandlers

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/dghubble/oauth1"
	otumblr "github.com/dghubble/oauth1/tumblr"
	tumblr "github.com/tumblr/tumblrclient.go"
)

type MyTumblrHandler struct {
	Tokens *APITokens
	Client *tumblr.Client
}

func InitHandler(configPath string) *MyTumblrHandler {
	log.Tracef("enter Init Handler with configpath: `%s`", configPath)
	t := &MyTumblrHandler{}
	t.Init(configPath)
	return t
}

func (t *MyTumblrHandler) Init(configPath string) {
	t.Tokens = &APITokens{}
	t.CreateClient()
	_, err := t.Client.GetDashboard()
	if err != nil {
		log.Fatal("could not initialize; failed to retrieve dashboard with error: " + err.Error())
	}
}

// To get consumer secret and consumer key, follow instructions here https://www.tumblr.com/oauth/apps
func (t *MyTumblrHandler) CreateClient() {
	if t == nil {
		log.Fatal("CreateClient: handler not initialized")
	}
	if t.Tokens.ConsumerKey == "" {
		log.Fatal("consumer key has no value")
	}
	if t.Tokens.ConsumerSecret == "" {
		log.Fatal("consumer secret has no value")
	}
	t.Client = tumblr.NewClientWithToken(t.Tokens.ConsumerKey, t.Tokens.ConsumerSecret, t.Tokens.AccessToken, t.Tokens.AccessSecret)
}

// Run this first by itself
func GetAccessToken(configPath string) {
	// https://github.com/dghubble/oauth1
	// https://github.com/dghubble/oauth1/blob/main/examples/tumblr-login.go
	t := &MyTumblrHandler{}
	t.Tokens = &APITokens{}
	t.Tokens.LoadFromJSON(configPath)
	config := oauth1.Config{
		ConsumerKey:    t.Tokens.ConsumerKey,
		ConsumerSecret: t.Tokens.ConsumerSecret,
		CallbackURL:    t.Tokens.CallBackURL,
		Endpoint:       otumblr.Endpoint,
	}
	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Fatal("failed to get request token " + err.Error())
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		log.Fatal("failed to authorize token " + err.Error())
	}

	log.Infof("Open this URL in your browser:\n\t%s\n", authorizationURL.String())

	log.Infof("Choose whether to grant the application access.\nPaste " +
		"the oauth_verifier parameter (excluding trailing #_=_) from the " +
		"address bar\n")

	splits := strings.Split(authorizationURL.String(), "authorize?oauth_token=")
	authUrl := splits[len(splits)-1]

	// New tokens and secrets; clear out deprecated content
	t.Tokens.OAuthToken = authUrl
	t.Tokens.RequestSecret = requestSecret
	t.Tokens.Verifier = ""
	t.Tokens.AccessSecret = ""
	t.Tokens.AccessToken = ""
	t.Tokens.SaveToJSON(configPath)
}

// Run this second
func GetOAuthToken(configPath string) {
	t := &MyTumblrHandler{}
	t.Tokens = &APITokens{}
	t.Tokens.LoadFromJSON(configPath)
	config := oauth1.Config{
		ConsumerKey:    t.Tokens.ConsumerKey,
		ConsumerSecret: t.Tokens.ConsumerSecret,
		CallbackURL:    t.Tokens.CallBackURL,
		Endpoint:       otumblr.Endpoint,
	}
	accessToken, accessSecret, err := config.AccessToken(t.Tokens.OAuthToken, t.Tokens.RequestSecret, t.Tokens.Verifier)
	if err != nil {
		log.Fatal("failed to get access token " + err.Error())
	}
	accessOAuthToken := oauth1.NewToken(accessToken, accessSecret)
	log.Infoln("Consumer was granted an access token to act on behalf of a user.")
	t.Tokens.AccessToken = accessOAuthToken.Token
	t.Tokens.AccessSecret = accessOAuthToken.TokenSecret
	t.Tokens.SaveToJSON(configPath)
}
