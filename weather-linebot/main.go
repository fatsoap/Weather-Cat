package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	env := GetEnv()
	bot, err := linebot.New(env.LINE_SCRET_TOKEN, env.LINE_ACCESS_TOKEN)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "LineBot SDK Init Failed.",
			StatusCode: 500,
		}, nil
	}
	var r http.Request
	r.Body = io.NopCloser(strings.NewReader(request.Body))
	r.Header = make(http.Header)
	r.Header.Add("X-Line-Signature", request.Headers["X-Line-Signature"])

	message_events, err := bot.ParseRequest(&r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			fmt.Println("ErrInvalidSignature")
		} else {
			fmt.Println("Bad Requst")
			fmt.Println(err)
		}
	}
	for _, event := range message_events {
		fmt.Println(event)
	}
	// 	if event.Type == linebot.EventTypeFollow { //加好友

	// 	} else if event.Type == linebot.EventTypeFollow { //加好友
	// 		//ts := event.Timestamp.Format("MM-dd-HH-mm")
	// 		//fmt.Println(event.Source.Type) // group , user, room
	// 		switch message := event.Message.(type) {
	// 		case *linebot.TextMessage: //是文字訊息
	// 			quota, err := s.bot.GetMessageQuota().Do()
	// 			if err != nil {
	// 				fmt.Println("Get Quota err:", err)
	// 			}
	// 			resMessage := MessageParser(message, quota)
	// 			if _, err = s.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resMessage)).Do(); err != nil {
	// 				fmt.Println("Line Bot Reply Message error : ", err)
	// 			}
	// 		default:
	// 			fmt.Println("Not text message") //不是文字訊息
	// 			if _, err = s.bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("6136", "10551376")).Do(); err != nil {
	// 				fmt.Println("Line Bot Reply Sticker error : ", err)
	// 			}
	// 			//list of sticker
	// 			//https://developers.line.biz/en/docs/messaging-api/sticker-list/#sticker-definitions
	// 		}
	// 	} else {
	// 		// 不需要回覆
	// 	}
	// }
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string("Sheep")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
