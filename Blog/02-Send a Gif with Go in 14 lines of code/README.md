## 02-Send-a-gif-with-Go-using-MMS-in-14-lines

In past posts we’ve shown you how to use Go to send SMS and how to make and receive calls with Twilio. What if I wanted to share a gopher gif with a friend using Go? Twilio can do that too!

```golang
package main

import (
  "fmt"
  twilio "github.com/kevinburke/twilio-go"
  "net/url"
)

func main() {
  client := twilio.NewClient("TWILIO_AUTH_SID", "TWILIO_AUTH_TOKEN", nil)
  gif, _ := url.Parse("https://media.giphy.com/media/uGGT9wVlxPAuk/giphy.gif")
  mms, _ := client.Messages.SendMessage("FROM_NUMBER", "TO_NUMBER", "", []*url.URL{gif})
  fmt.Println("You just sent a gif with Twilio using Go!! " + string(mms.Status) + " - " + string(mms.Sid))
}
```

### Setting Up

To start you will need a few things:

* Twilio Account (sign up here)
* Twilio Phone Number
* Go

Once you have completed installing Go, install the twilio-go package with the get command.

```golang
go get github.com/kevinburke/twilio-go
```

After completing the installation create a new Go program. Import the three libraries below: the twilio-go library from github and the Go native libraries net/url and fmt. Complete the template by adding the program’s main function.

```golang
package main

import (
     "fmt"
     twilio "github.com/kevinburke/twilio-go"
     "net/url"
)

func main(){
// code goes here
}
```

### Sending a MMS with Twilio

Next create a new variable called client. Initialize your client using twilio.NewClient and pass in your Account SID and Auth Token. These credentials can be found on the main page of the Twilio Console.

```golang
client := twilio.NewClient("TWILIO_AUTH_SID", "TWILIO_AUTH_TOKEN", nil)
```

It's ok to hard code credentials while you're getting started, but make sure you don't share this file or commit these to GitHub. Now find a gif you would like to share. For this example I selected a gopher because it is the mascot of Go.

Once you’ve found a gif to share, parse the file so it can be passed to the API using twilio-go. Declare a new variable called gif and instantiate the url.Parse method from the net/url library.

```golang
gif, _ := url.Parse("https://media.giphy.com/media/uGGT9wVlxPAuk/giphy.gif")
```

The url.Parse method returns a pointer to the string url in memory and an error value by default. You may notice the “_” or blank identifier that is next to the declared variable gif. Using the blank identifier specifies that the second value, the error value in this case, will be ignored.

Now to the gopher stuff, let’s send a gif. Create a new variable called mms and initialize your message using client.Messages.SendMessage.

```golang
mms, _ := client.Messages.SendMessage("FROM_NUMBER", "TO_NUMBER", "", []*url.URL{gif})
```

A few pieces of information are needed for the SendMessage method to make this happen.

* **FROM_NUMBER:** Your Twilio Number
* **TO_NUMBER:** Number that is receiving the gif
* **"":** Empty text body - we're sending a picture so we don't need this! 
  * If you want to add a body to the message you can write the message in quotations. (i.e. “How do you like the gopher?”)
* __[]*url.URL{gif}:__ the MMS (picture) body of the message

Lastly let’s print out a message, the MMS status and the MMS SID.

```golang
fmt.Println("You just sent a gif with Twilio using Go! " + string(mms.Status) + " - " + string(mms.Sid))
```

Ready? Let’s Go! Run your program from the command line with the following:

```golang
go run Send-a-Gif-with-Go-in-14-lines-of-code.go
```

### Complete Program

```golang
package main

import (
     "fmt"
     twilio "github.com/kevinburke/twilio-go"
     "net/url"
)

func main() {
     client := twilio.NewClient("TWILIO_AUTH_SID", "TWILIO_AUTH_TOKEN", nil)
     gif, _ := url.Parse("https://media.giphy.com/media/uGGT9wVlxPAuk/giphy.gif")
     mms, _ := client.Messages.SendMessage("FROM_NUMBER", "TO_NUMBER", "", []*url.URL{gif})
     fmt.Println("You just sent a gif with Twilio using Go!! " + string(mms.Status) + " - " + string(mms.Sid))
}
```

### What now?

You just found out how easy it is to send a gif using Twilio with Go and the fun is just beginning. Try integrating this idea into an existing app or building an app using concurrency with Go’s goroutines. The sky’s the limit.

If you have any questions about how to use Twilio & Go drop me a line anytime via one of the mediums below. This project can be found on my GitHub in the repository called TwilioGodex, a rolodex of Twilio based Go projects. 

* Email: ckonopka@twilio.com
* Twitter: @cskonopka
* Github: cskonopka
