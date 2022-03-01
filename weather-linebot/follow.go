package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

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

func handleFollow(event *linebot.Event) error {
	return nil
}
