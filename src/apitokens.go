package mytumblrhandlers

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type APITokens struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	CallBackURL    string `json:"callBackURL"`
	Verifier       string `json:"verifier"`
	OAuthToken     string `json:"oauthToken"`
	RequestSecret  string `json:"requestSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

func (tokens *APITokens) SaveToJSON(path string) (ok bool) {
	log.Trace("saving to json")
	apitokens, err := json.MarshalIndent(*tokens, "", "  ")
	if err != nil {
		log.Warning("could not save apitoken with error: " + err.Error())
		return
	}
	err = os.WriteFile(path, apitokens, 0755)
	if err != nil {
		log.Warning("could not write apitoken with error: " + err.Error())
		return
	}
	return true
}

func (tokens *APITokens) LoadFromJSON(path string) {
	log.Trace("loading from json")
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("could not read JSON with consumerSecret or consumerKey in path with error: " + err.Error())
	}

	err = json.Unmarshal(dat, tokens)
	if err != nil {
		log.Fatal("could not unmarshal JSON with consumerSecret or consumerKey with error: " + err.Error())
	}
}
