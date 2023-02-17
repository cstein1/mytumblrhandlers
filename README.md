# mytumblrhandlers
## Intro
This is a personal repo for accessing Tumblr in an easier fashion.

## To fill out config.secret
See `config.secret` in this project.
Please fill the keys `consumerKey`, `consumerSecret`, and `callBackURL`. Register a Tumblr app [here](https://www.tumblr.com/oauth/apps) to find the Consumer Key and Consumer Secret, as well as registering a Callback URL.

### Steps to fill out required fields
- Fill `consumerKey`, `consumerSecret`, and `callBackURL` in the config file
- Either call `GetAccessToken(configSecretPath)` with a user-provided `configSecretPath`, or call `go run get_tokens.go` with `cfg/config.secret` filled out, and see the following output
>Open this URL in your browser:
>https://www.tumblr.com/oauth/authorize?oauth_token=CbAzYxOhMyWhatAStrangeTokenxYzAbC
>Choose whether to grant the application access.
>Paste the oauth_verifier parameter (excluding trailing `#_=_`) from the address bar
- Grab that URI printed above and...
  - Enter it into your favorite browser
  - After allowing your app, the browser will add a `?oauth_verifier` parameter in the URI which you must copy and paste into the config under the `verifier` key (without the trailing `#_=_`)
  - The the `oauthToken`, `requestSecret` variables inside of the config will be automatically populated
- Everything except the `accessToken` and `accessSecret` are now populated, so time to get those!
- Run `GetOAuthToken` or edit the var inside __get_tokens.go__ called `FIRSTRUN` to `false` and then run `go run get_tokens.go` again to get authorization; now every key will be populated!
>Consumer was granted an access token to act on behalf of a user.

**Congratulations!** you should now have access to the Tumblr client!! See `src/initialization_test.go` for an example call after the `config.secret` has been properly populated.

### Post Summary Example:
The following is a minimal example of retrieving summaries of text posts.
```
import (
	"fmt"
	"path/filepath"
	"strings"
	mth "github.com/cstein1/mytumblrhandlers/src"
)

func main() {
  here, _ := filepath.Abs(".")
  DEFAULTCONFIGLOCATION = filepath.Join(here, "cfg", "config.secret")
  ok := mth.InitHandler(DEFAULTCONFIGLOCATION)
  if !ok {
    return
  }
  blogObj, err := mth.GetBlog("staff")
  if err != nil {
    return
  }
  posts, lastPost, _ := mth.GetTextPostsSummary(blogObj, mth.NOWTIME, 2)
  fmt.Println(posts)
  posts, _, _ = mth.GetTextPostsSummary(blogObj, lastPost, 2)
  fmt.Println(posts)
}
```
Which results in:
```
Twitter’s API changes***Ok folks. Everybody stay calm. It’s happening. Polls are here.
🔔***Listen up, friends. You may have noticed that some little elves here at Tumblr HQ have decorated your dashboard with enchanted...
```
Using a Summary will not result in the entire post being returned, but it returns non-HTML, clean content.

### Get Thread Example

The following is a minimal example of retrieving threads from text posts. `GetTextPostThread` returns `posts`, `lastPost`, and `err`, which are a list of list of string. 
This is a list of text posts including responses. If this call returns 20 posts that are each a text post with no responses, then the `posts` return will be a list of length 20, where each entry is of length 1.
```
import (
  "fmt"
  "path/filepath"
  "strings"
  mth "github.com/cstein1/mytumblrhandlers/src"
)

func main() {
  here, _ := filepath.Abs(".")
  DEFAULTCONFIGLOCATION = filepath.Join(here, "cfg", "config.secret")
  ok := mth.InitHandler(DEFAULTCONFIGLOCATION)
  if !ok {
    return
  }
  blogObj, err := mth.GetBlog("staff")
  if err != nil {
    return
  }
  posts, lastPost, err := mth.GetTextPostThread(blogObj, mth.NOWTIME, 2)
  if err != nil {
    return
  }
  mth.PrettyPrintTrail(posts)
  nextPosts, _, err := mth.GetTextPostThread(blogObj, lastPost, 2)
  if err != nil {
    return
  }
  mth.PrettyPrintTrail(nextPosts)
}
```
Which results in
```
- staff: Twitter’s API changes
Hello, Tumblr. Twitter unexpectedly announced they will end free access to the Twitter API (Application Programming Interface) on February 9th. A 
recent update
 had extended that date to today, February 13th, 2023. 
What’s an API?
Put simply, APIs define how programming systems interact with one another. For example, an API allows third parties (say, a social media platform) to interact with another website programmatically (to publish posts, etc.)
How does this affect Tumblr?
We are removing the feature that 
links a Twitter account
 to your Tumblr. This will end the ability to automatically share your new posts directly
 from Tumblr to Twitter 
when they are published. This will also remove the ability to automatically 
display recent tweets in your blog’s theme.
 This change will immediately affect web, Android, and iOS.
We appreciate this will be frustrating for many users, but all is not lost—you 
will
still be able to share any Tumblr post to Twitter manually using the normal share option on web. For Android and iOS, you will need to have the Twitter app installed to manually share Tumblr posts on Twitter. This functionality is expected for both platforms in upcoming app versions.

===============
- staff: Ok folks. Everybody stay calm. It’s happening. Polls are here.
Do you love polls?
yes
of course
naturally
i want some toast
puss in boots: the last wish really got to me and i’m feeling a lot of feelings
See Results
We’re starting to roll out a polls feature for the post editor across all platforms. Yes, that’s
all platforms
. That’s iOS, Android, and web. You should all have access to this in the next few days. And what’s more, it’s super easy to use.
Simply select the newly added orange poll icon (on web, it appears when you start a new content block in the post editor, on mobile it’s in your toolbar).
Write your question.
Add at least 2, at most 10 response options for people to choose from. There’s a character limit on these, it’s about as long as the
Puss in Boots
 option offered above.
Set the poll duration to one day or one week. 
You can add some commentary to your poll or let it speak for itself. 
And, just as with any post, you can toggle who sees the poll by using your community labels. Or you can go wild and blaze it if it meets the content guidelines for blazing. 
Come across a poll in the wild? Have your say! Click on your chosen response. Congratulations, you’ve just voiced an opinion.
And that’s all there is to it. Have fun out there.

===============
- staff: 🔔
On December 24, some bells appeared on your dashboards on web at Tumblr dot com. At first, you were confused. Then, you were delighted.
The bells stayed for ~3 days. In that time:
You
jingled
 the bells
40 million times!
123,000 of you jingled the bells at least once. 
90,000 of you
jingled
 the bells long enough to unlock a flurry of snow. 
Nearly 2,400 of you discovered a hidden feature that allowed you to play tunes via the search bar…and used that 14,000 times to play tunes.
Many
 of you
created
 and
shared
 “sheet music”—strings of numbers that other users could press or paste into the search bar to play. Cute! 
The elves were extremely pleased with themselves. We get the sense that more elfish mischief might be coming our way in the nearish future. Well done, Tumblr.

===============
- staff: Listen up, friends. You may have noticed that some little elves here at Tumblr HQ have decorated your dashboard with enchanted bells. We told them not to, but elves are notoriously single-minded. (You might even say jingle-minded if you were so inclined, which we most certainly are not.)
So here’s the deal. You can find the bells on web on your desktop by clicking “Click for bells.” You can look at them. You can jangle them. Heck, you can even play whole songs on ’em if you so please. They’ll be up for a short time only. If you clicked for bells and hate them, refresh your page, and it will be blissfully bell-free.
Rumor has it that crabs enjoy bells, while elfish legend speaks of excessive tolling bringing little miracles down from above. This is all hearsay, though, and not to be believed (we all know how elves like to tell tall tales. You have been warned).
(thanks to
freesfx

for the dulcet chimes)

===============
```

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
