package mytumblrhandlers

import (
	"encoding/json"
	"os"
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

func (at *APITokens) SaveToJSON(path string) {
	apitokens, err := json.MarshalIndent(*at, "", "  ")
	if err != nil {
		panic("could not save apitoken")
	}
	err = os.WriteFile(path, apitokens, 0755)
	if err != nil {
		panic("could not write apitoken")
	}
}

func (tokens *APITokens) LoadFromJSON(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic("could not read JSON with consumerSecret or consumerKey in path " + path)
	}

	err = json.Unmarshal(dat, tokens)
	if err != nil {
		panic("could not unmarshal JSON with consumerSecret or consumerKey: " + err.Error())
	}
}
