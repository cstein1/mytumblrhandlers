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

var EMPTYCONFIG map[string]string = map[string]string{
	"consumerKey":    "",
	"consumerSecret": "",
	"callBackURL":    "",
	"verifier":       "",
	"oauthToken":     "",
	"requestSecret":  "",
	"accessToken":    "",
	"accessSecret":   "",
}

var input *bufio.Scanner
var configLocation string
var printCmdline bool
var config mth.APITokens

var repeat string = "Seems you already have a %s. Continuing...\n"

const (
	PANIC = "TERMINAL FAILURE. PLEASE TRY AGAIN. PROGRESS MAY BE SAVED. USE THE SAME LOCATION FOR YOUR CONFIG TO LOAD PROGRESS."
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	config = mth.APITokens{}
	welcomeMessage := `
	Welcome to MyTumblrHandlers! I will be guiding you through getting your API Tokens config set up!
	This application will build a config.secret file that mytumblrhandlers will pull from when calling the Tumblr API.
	If you have the information required to run this already, feel free to stop this application.
	At the end of this guide, you'll be able to edit the components manually via vim/notepad/etc... so don't panic if you get something wrong.
	All you need is a config with the following keys/values included:
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

	Feel free to copy this, save it somewhere, and pass it to your mytumblrhandler init call!
	Otherwise, this program will help guide you to fill in these components!
	`
	fmt.Println(welcomeMessage)
	prettyEnter()
	input.Scan()
	fmt.Println("Okay, let's do this! Please start by registering your Tumblr App here: https://www.tumblr.com/oauth/apps")
	enterCallBackURL()
	fmt.Println("I will be receiving text from you via commandline. I will also be sending you information via commandline. So let's do this!")
	handleLocation()
	fmt.Println("Okay, now that we got that out of the way, let's grab those Consumer components from you.")
	enterConsumerKey()
	enterConsumerSecret()
	getAccessTokenPrompts()
	mth.GetOAuthToken(configLocation)
	fmt.Println("Hit enter to try to get your dashboard. If this works then we're done here and you can use the mytumblrhandlers in your code with the created config!")
	input.Scan()
	dash, err := mth.GetDashboard()
	if err != nil {
		fmt.Println("It didn't work! Couldn't get to your dash :(")
		panic(PANIC)
	}
	fmt.Printf("%v", dash)
	fmt.Println("Congratulations! It worked! Find your config in %s", configLocation)
}

func getAccessTokenPrompts() {
	var ok bool
	authorizationUrl, ok := mth.GetAccessToken(configLocation)
	if !ok {
		fmt.Println("Failed to get access tokens!")
		panic(PANIC)
	}
	msg := `
	Please go ahead and open this URL in your favorite browser %s 
	This URL will bring you Tumblr where you can give your app permission to access your tumblr.
	THIS IS IMPORTANT. When you grant access, you will be redirected to your dashboard.
	BEFORE YOU DO ANYTHING ELSE: copy and paste your URL here and hit enter.
	`
	fmt.Printf("\n"+msg+"\n", authorizationUrl)
	input.Scan()
	txt := string(input.Text())
	u, err := url.Parse(txt)
	if err != nil {
		fmt.Println("Failed to parse with error: " + err.Error())
		panic(PANIC)
	}
	params := u.Query()
	config.OAuthToken = params.Get("oauth_verifier")
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
		input.Scan()
		txt := string(input.Text())
		if len(txt) <= 1 {
			fmt.Println("Please make sure the text is longer than just the newline character! :)")
			continue
		}
		config.CallBackURL = strings.Trim(txt, " \n\t")
		break
	}
}

func enterConsumerKey() {
	if config.ConsumerKey != "" {
		fmt.Printf(repeat, "consumerkey")
		return
	}
	fmt.Println("Please enter your consumer key; you should be able to find it on the summary page for your applications")
	for {
		input.Scan()
		txt := string(input.Text())
		if len(txt) <= 1 {
			fmt.Println("Please make sure the text is longer than just the newline character! :)")
			continue
		}
		config.ConsumerKey = strings.Trim(txt, " \n\t")
		break
	}
}

func enterConsumerSecret() {
	if config.ConsumerSecret != "" {
		fmt.Printf(repeat, "consumersecret")
		return
	}
	fmt.Println("Please enter your consumer key; you should be able to find it on the summary page for your applications")
	for {
		input.Scan()
		txt := string(input.Text())
		if len(txt) <= 1 {
			fmt.Println("Please make sure the text is longer than just the newline character! :)")
			continue
		}
		config.ConsumerSecret = strings.Trim(txt, " \n\t")
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

func handleLocation() (configLocation string) {
	for {
		fmt.Println("Would like me to store your config in the (D)efault spot in ./cfg/config.secret or if you would like to specify a (C)ustom place to save this config to: (D/c)")
		txt := string(input.Text())
		if len(txt) > 2 {
			fmt.Println("Please only select one letter or hit the return/enter key!")
			input.Scan()
			continue
		}
		switch string(strings.ToLower(txt)[0]) {
		case "d", "\n", "":
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
				config.LoadFromJSON(configLocation)
			}
			break
		case "c":
			fmt.Println("Please select where you would like your file to be placed: ")
			input.Scan()
			givenPath := input.Text()
			realPath, err := filepath.Abs(givenPath)
			if err != nil {
				fmt.Printf("\nFAILED TO ACCESS THAT LOCATION. PLEASE TRY AGAIN. SEE ERROR: %s", err.Error())
				continue
			}
			if _, err := os.Stat(realPath); os.IsNotExist(err) {
				dir := filepath.Dir(realPath)
				err = os.MkdirAll(dir, mth.USR_RW)
				if err != nil {
					fmt.Printf("\nFAILED TO CREATE A DIR FOR LOCATION. PLEASE TRY AGAIN. SEE ERROR: %s", err.Error())
					continue
				}
			} else {
				// if it does exist, this might have been done already. Let's load the JSON
				config.LoadFromJSON(realPath)
			}
			configLocation = realPath
			break
		default:
			fmt.Println("Did not recognize your response; please try again")
		}
	}
}
