package main

// This application will build the config required to use mytumblrhandlers

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	mth "github.com/cstein1/mytumblrhandlers/src"
)

var input *bufio.Scanner
var configLocation string
var printCmdline bool
var config mth.APITokens
var cmdlineText string
var superSecretSkip bool

var repeat string = "Seems you already have a %s. Continuing...\n"

const (
	PANIC = "TERMINAL FAILURE. PLEASE TRY AGAIN. PROGRESS MAY BE SAVED. USE THE SAME LOCATION FOR YOUR CONFIG TO LOAD PROGRESS."
)

func main() {
	config = mth.APITokens{}
	welcomeMessage := `
	Welcome to MyTumblrHandlers! I will be guiding you through getting your API Tokens config set up!
	This application will build a config.secret file that mytumblrhandlers will pull from when calling the Tumblr API.
	If you have the information required to run this already, feel free to stop this application and use that config in the mytumblrhandlers.Init call.
	At the end of this guide, you'll be able to edit the components manually via vim/notepad/etc... so don't panic if you get something wrong.
	At the end of this program, will have a config with the following keys/values included:
	{
		"consumerKey":    "",
		"consumerSecret": "",
		"callBackURL":    "",
		"verifier":       "",
		"oauthToken":     "",
		"requestSecret":  "",
		"accessToken":    "",
		"accessSecret":   ""
	}

	At the end of the day, we only need tokens.ConsumerKey, tokens.ConsumerSecret, tokens.AccessToken, tokens.AccessSecret.
	If you have those, feel free to copy this JSON, fill in the values you want, save it somewhere, and pass it to your mytumblrhandler init call!
	If you don't have that info, or don't know where to get it, then this program will help guide you to fill in these components!
	`
	fmt.Println(welcomeMessage)
	prettyEnter()
	readLine()
	fmt.Println("Okay, let's do this! Please start by registering your Tumblr App here: https://www.tumblr.com/oauth/apps and make sure to write down your callback url for later!")
	superSecretSkip = handleLocation()
	if !superSecretSkip {
		enterCallBackURL()
		enterConsumerKey()
		enterConsumerSecret()
		getAccessTokenPrompts()
		mth.GetOAuthToken(configLocation)
		fmt.Println("Trying to get a couple of posts from Staff. If this works then we're done here and you can use the mytumblrhandlers in your code with the created config!")
	}
	mth.Init(configLocation)
	blogObj, err := mth.GetBlog("staff")
	if err != nil {
		fmt.Println("It didn't work! Couldn't staff blog!: " + err.Error())
		panic(PANIC)
	}
	posts, _, _ := mth.GetTextPostsSummary(blogObj, mth.NOWTIME, 10)
	fmt.Println(posts)
	fmt.Println("Congratulations! It worked! Find your config in " + configLocation)
}

func readLine() string {
	cmdlineText = ""
	fmt.Scanln(&cmdlineText)
	return cmdlineText
}

func getAccessTokenPrompts() {
	var ok bool
	fmt.Println("The next call might take a second; please be patient. Don't buffer any keypresses.")
	fmt.Printf("Getting the access tokens to %s ", configLocation)
	authorizationUrl, ok := mth.GetAccessToken(configLocation)
	if !ok {
		fmt.Println("Failed to get access tokens!")
		panic(PANIC)
	}
	// GetAccessToken saves the JSON which we don't want to overwrite with our current config value
	config.LoadFromJSON(configLocation)
	msg := `
	Please go ahead and open this URL in your favorite browser %s 
	This URL will bring you Tumblr where you can give your app permission to access your tumblr.
	THIS IS IMPORTANT. When you grant access, you will be redirected to your dashboard.
	BEFORE YOU DO ANYTHING ELSE: copy and paste that new redirected URL here and hit enter.\n
	`
	fmt.Printf(msg+":", authorizationUrl)
	for {
		readLine()
		if cmdlineText == "" {
			fmt.Println("Clearing buffer. Please try again")
			fmt.Print("Enter the redirected URL here: ")
			continue
		}
		break
	}
	u, err := url.Parse(cmdlineText)
	if err != nil {
		fmt.Println("Failed to parse with error: " + err.Error())
		panic(PANIC)
	}
	params := u.Query()
	config.Verifier = params.Get("oauth_verifier")
	config.OAuthToken = params.Get("oauth_token")
	ok = config.SaveToJSON(configLocation)
	if !ok {
		fmt.Println("Failed to save oauth verifier!")
		panic(PANIC)
	}
}

func enterCallBackURL() {
	if config.CallBackURL != "" {
		fmt.Printf(repeat, "callbackurl")
		return
	}
	for {
		fmt.Println("While you're filling in your application, you'll be asked for a callback url! My callback url is the URL to my bot, but it can be anything.")
		fmt.Print("Please enter that callback url here: ")
		readLine()
		if len(cmdlineText) <= 1 {
			fmt.Println("Please make sure the text is longer than just the newline character! :)")
			continue
		}
		config.CallBackURL = strings.Trim(cmdlineText, " \n\t")
		break
	}
}

func enterConsumerKey() {
	if config.ConsumerKey != "" {
		fmt.Printf(repeat, "consumerkey")
		return
	}
	fmt.Print("Please enter your consumer key; you should be able to find it on the summary page for your applications: ")
	for {
		readLine()
		if len(cmdlineText) <= 1 {
			fmt.Println("Please make sure the text is longer than just the newline character! :)")
			continue
		}
		config.ConsumerKey = strings.Trim(cmdlineText, " \n\t")
		break
	}
}

func enterConsumerSecret() {
	if config.ConsumerSecret != "" {
		fmt.Printf(repeat, "consumersecret")
		return
	}
	fmt.Print("Please enter your consumer key; you should be able to find it on the summary page for your applications: ")
	for {
		readLine()
		if len(cmdlineText) <= 1 {
			fmt.Println("Please make sure the text is longer than just the newline character! :)")
			continue
		}
		config.ConsumerSecret = strings.Trim(cmdlineText, " \n\t")
		break
	}
}

func prettyEnter() {
	fmt.Print("\nHit enter to continue")
	sleepTime := time.Duration(400)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * sleepTime)
		fmt.Print(".")
	}
}

func handleLocation() (secretSkip bool) {
	for {
		fmt.Print("Would you like me to store your config in the (D)efault spot in ./cfg/config.secret or if you would like to specify a (C)ustom place to save this config to: (D/c) ")
		readLine()
		switch string(cmdlineText) {
		case "d", "":
			here, err := filepath.Abs(".")
			if err != nil {
				fmt.Println("FAILED TO FIND HERE: " + err.Error())
				continue
			}
			dirLocation := filepath.Join(here, "cfg")
			if _, err := os.Stat(dirLocation); os.IsNotExist(err) {
				err = os.Mkdir(dirLocation, mth.USR_RW)
				if err != nil {
					fmt.Println("FAILED TO MAKE DIR " + dirLocation + " WITH ERR: " + err.Error())
					continue
				}
			}
			configLocation = filepath.Join(dirLocation, "config.secret")
			if _, err := os.Stat(configLocation); !os.IsNotExist(err) {
				if ok := config.LoadFromJSON(configLocation); !ok {
					panic(PANIC)
				}
			}
			return
		case "c":
			fmt.Println("Please select where you would like your file to be placed: ")
			readLine()
			givenPath := cmdlineText
			realPath, err := filepath.Abs(givenPath)
			if err != nil {
				fmt.Printf("\nFAILED TO ACCESS THAT LOCATION. PLEASE TRY AGAIN. SEE ERROR: %s\n", err.Error())
				continue
			}
			if _, err := os.Stat(realPath); os.IsNotExist(err) {
				dir := filepath.Dir(realPath)
				err = os.MkdirAll(dir, mth.USR_RW)
				if err != nil {
					fmt.Printf("\nFAILED TO CREATE A DIR FOR LOCATION. PLEASE TRY AGAIN. SEE ERROR: %s\n", err.Error())
					continue
				}
			} else {
				// if it does exist, this might have been done already. Let's load the JSON
				if ok := config.LoadFromJSON(realPath); !ok {
					panic(PANIC)
				}
			}
			configLocation = realPath
			return
		case "o":
			fmt.Println("Secret, fast, unsafe skip activated. You already have a perfect config and just want to test it! I see you :)")
			here, _ := filepath.Abs(".")
			dirLocation := filepath.Join(here, "cfg")
			configLocation = filepath.Join(dirLocation, "config.secret")
			if ok := config.LoadFromJSON(configLocation); !ok {
				fmt.Println("Hope you didn't mean to enter the super secret skip!")
				panic(PANIC)
			}
			return true
		default:
			fmt.Printf("\nDid not recognize your response of %s; please try again\n", cmdlineText)
		}
	}
}
