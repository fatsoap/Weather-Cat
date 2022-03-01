package main

import (
	"fmt"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

/*
Follow Event Struct
{
  "destination": "xxxxxxxxxx",
  "events": [
    {
      "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
      "type": "follow",
      "mode": "active",
      "timestamp": 1462629479859,
      "source": {
        "type": "user",
        "userId": "U4af4980629..."
      }
    }
  ]
}
*/

func handleFollow(bot *linebot.Client, event *linebot.Event) error {
	fmt.Printf("User:%v\nFollowed Weather Cat", event.Source.UserID)
	text := fmt.Sprintf(hintFlexMessageTemplate, fmt.Sprintf("%v", time.Now().Nanosecond())[:2])
	altText := defaultAltText
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(text))
	if err != nil {
		fmt.Println("Flex Message Unmarshal Failed")
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage(altText, flexContainer)).Do(); err != nil {
		fmt.Println("Reply Flex Message Failed")
	}
	return nil
}
