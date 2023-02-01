# mytumblrhandlers
## Intro
This is a personal repo for accessing Tumblr in an easier fashion. 

## To fill out config.secret
In the top layer of this project is a file called `config.secret`.
Please fill in the `consumerKey`, `consumerSecret`, and `callBackURL`. Register a Tumblr app [here](https://www.tumblr.com/oauth/apps) to find the Consumer Key and Consumer Secret, as well as registering a Callback URL.

### Steps to fill out required fields
- Fill `consumerKey`, `consumerSecret`, and `callBackURL` in the config file
- Either call `GetAccessToken(configSecretPath)` with a user-provided `configSecretPath`, or call `go run get_tokens.go` with `src/config.secret` filled out, and see the following output
>Open this URL in your browser:
>https://www.tumblr.com/oauth/authorize?oauth_token=CbAzYxOhMyWhatAStrangeTokenxYzAbC
>Choose whether to grant the application access.
>Paste the oauth_verifier parameter (excluding trailing #_=_) from the address bar
- - Take the authorize URI printed to the console, and enter it into your favorite browser
- - After allowing your app, the browser will add a `?oauth_verifier` parameter in the URI which you must copy and paste into the config under the `verifier` key (without the trailing `#_=_`)
- - The the `oauthToken`, `requestSecret` variables inside of the config will be automatically populated
- Everything except the `accessToken` and `accessSecret` are now populated, so time to get those!
- Run `GetOAuthToken` or edit the var inside __get_tokens.go__ called `FIRSTRUN` to `false` and then run `go run get_tokens.go` again to get authorization; now every key will be populated!
>Consumer was granted an access token to act on behalf of a user.

**Congratulations!** you should now have access to the Tumblr client!! See `src/initialization.go` for an example call after the config.secret has been properly populated.


## FAQ: 
If something doesn't work, try getting the following:
Installing packages that aren't included in the go.mod, but are required 
```
sudo apt-get update
sudo apt-get install build-essential
sudo apt-get install gcc
```

## Resources
https://github.com/dghubble/oauth1
https://github.com/dghubble/oauth1/blob/main/examples/tumblr-login.go