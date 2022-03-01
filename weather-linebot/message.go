package main

import (
	"fmt"

	owm "github.com/briandowns/openweathermap"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

/*
{
  "destination": "xxxxxxxxxx",
  "events": [
    {
      "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
      "type": "message",
      "mode": "active",
      "timestamp": 1462629479859,
      "source": {
        "type": "user",
        "userId": "U4af4980629..."
      },
      "message": {
        "id": "325708",
        "type": "location",
        "title": "my location",
        "address": "Japan, 〒160-0004 Tokyo, Shinjuku City, Yotsuya, 1-chōme-6-1",
        "latitude": 35.687574,
        "longitude": 139.72922
      }
    }
  ]
}
*/

func handleMessage(event *linebot.Event) error {
	switch message := event.Message.(type) {
	case *linebot.LocationMessage: // 位置訊息
		lat := message.Latitude
		lon := message.Longitude
		fmt.Println(lat, lon) // TODO: call open weather api
	default: // 其他訊息
		fmt.Println("Not Location Message")
	}

	return nil
}
