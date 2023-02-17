package mytumblrhandlers

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/dghubble/oauth1"
	otumblr "github.com/dghubble/oauth1/tumblr"
	tumblr "github.com/tumblr/tumblrclient.go"
)

var tokens *APITokens
var client *tumblr.Client

func InitHandler(configPath string) (ok bool) {
	log.Tracef("enter Init Handler with configpath: `%s`", configPath)
	return Init(configPath)
}

func Init(configPath string) (ok bool) {
	tokens = &APITokens{}
	tokens.LoadFromJSON(configPath)
	if tokens.ConsumerKey == "" || tokens.ConsumerSecret == "" {
		log.Fatal("config does not provide consumer key or consumer secret")
	}
	return CreateClient()
}

// To get consumer secret and consumer key, follow instructions here https://www.tumblr.com/oauth/apps
func CreateClient() (ok bool) {
	if tokens == nil {
		log.Fatalf("tokens object not initialized. Please see cfg/config.secret or create your own config file.")
		return
	}
	if tokens.ConsumerKey == "" {
		log.Fatal("consumer key has no value")
	}
	if tokens.ConsumerSecret == "" {
		log.Fatal("consumer secret has no value")
	}
	client = tumblr.NewClientWithToken(tokens.ConsumerKey, tokens.ConsumerSecret, tokens.AccessToken, tokens.AccessSecret)
	return client != nil
}

// Run this first by itself
func GetAccessToken(configPath string) (url string, ok bool) {
	// https://github.com/dghubble/oauth1
	// https://github.com/dghubble/oauth1/blob/main/examples/tumblr-login.go
	tokens = &APITokens{}
	tokens.LoadFromJSON(configPath)
	config := oauth1.Config{
		ConsumerKey:    tokens.ConsumerKey,
		ConsumerSecret: tokens.ConsumerSecret,
		CallbackURL:    tokens.CallBackURL,
		Endpoint:       otumblr.Endpoint,
	}
	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Warning("failed to get request token " + err.Error())
		return
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		log.Warning("failed to authorize token " + err.Error())
		return
	}

	// log.Infof("Open this URL in your browser:\n\t%s\n", authorizationURL.String())

	// log.Infof("Choose whether to grant the application access.\nPaste " +
	// 	"the oauth_verifier parameter (excluding trailing #_=_) from the " +
	// 	"address bar\n")

	splits := strings.Split(authorizationURL.String(), "authorize?oauth_token=")
	authUrl := splits[len(splits)-1]

	// New tokens and secrets; clear out deprecated content
	tokens.OAuthToken = authUrl
	tokens.RequestSecret = requestSecret
	tokens.Verifier = ""
	tokens.AccessSecret = ""
	tokens.AccessToken = ""
	return authorizationURL.String(), tokens.SaveToJSON(configPath)
}

// Run this second
func GetOAuthToken(configPath string) (ok bool) {
	tokens = &APITokens{}
	tokens.LoadFromJSON(configPath)
	config := oauth1.Config{
		ConsumerKey:    tokens.ConsumerKey,
		ConsumerSecret: tokens.ConsumerSecret,
		CallbackURL:    tokens.CallBackURL,
		Endpoint:       otumblr.Endpoint,
	}
	accessToken, accessSecret, err := config.AccessToken(tokens.OAuthToken, tokens.RequestSecret, tokens.Verifier)
	if err != nil {
		log.Fatal("failed to get access token " + err.Error())
		return
	}
	accessOAuthToken := oauth1.NewToken(accessToken, accessSecret)
	// log.Infoln("Consumer was granted an access token to act on behalf of a user.")
	tokens.AccessToken = accessOAuthToken.Token
	tokens.AccessSecret = accessOAuthToken.TokenSecret
	return tokens.SaveToJSON(configPath)
}
