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
	fmt.Println("You just sent a gopher gif with Go! " + string(mms.Status) + " - " + string(mms.Sid))
}
