package robbot

import (
	"ez_line_bot/weather"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Server struct {
	*mux.Router
	bot *linebot.Client
}

func Init() *Server {
	bot, err := linebot.New(os.Getenv("LINE_SCRET_TOLEN"), os.Getenv("LINE_ACCESS_TOKEN"))
	if err != nil {
		fmt.Printf("Linebot init fail, message : %v\n", err)
	}
	s := &Server{
		Router: mux.NewRouter(),
		bot:    bot,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/bot", s.botRouter()).Methods("GET")
	s.HandleFunc("/bot", s.botRouter()).Methods("POST")
}

func (s *Server) botRouter() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		events, err := s.bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				fmt.Println("Server Error, Line Bot fail, meesage : ", err)
				http.Error(res, "Line Bot fail", http.StatusInternalServerError)
			} else {
				fmt.Println("Bad Request, Line Bot fail, meesage : ", err)
				http.Error(res, "Line Bot fail", http.StatusBadRequest)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage { //是訊息
				//ts := event.Timestamp.Format("MM-dd-HH-mm")
				//fmt.Println(event.Source.Type) // group , user, room
				switch message := event.Message.(type) {
				case *linebot.TextMessage: //是文字訊息
					quota, err := s.bot.GetMessageQuota().Do()
					if err != nil {
						fmt.Println("Get Quota err:", err)
					}
					resMessage := MessageParser(message, quota)
					if _, err = s.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resMessage)).Do(); err != nil {
						fmt.Println("Line Bot Reply Message error : ", err)
					}
				default:
					fmt.Println("Not text message") //不是文字訊息
					if _, err = s.bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("6136", "10551376")).Do(); err != nil {
						fmt.Println("Line Bot Reply Sticker error : ", err)
					}
					//list of sticker
					//https://developers.line.biz/en/docs/messaging-api/sticker-list/#sticker-definitions
				}
			} else {
				fmt.Println("Not message") //不是訊息
				if _, err = s.bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("6359", "11069850")).Do(); err != nil {
					fmt.Println("Line Bot Reply Sticker error : ", err)
				}
			}
		}
	}
}

func quotaParser(quota *linebot.MessageQuotaResponse) string {
	q := "配額"
	q += "\n使用配額:" + strconv.FormatInt(quota.TotalUsage, 10)
	q += "\n剩餘配額:" + strconv.FormatInt(quota.Value, 10)
	q += "\n目前方案:" + quota.Type
	return q
}

func MessageParser(msg *linebot.TextMessage, quota *linebot.MessageQuotaResponse) string {

	command := "指令\n"
	command += "帥氣指數\n"
	command += "天氣台北 (台北、台中、台南、中壢)\n"
	command += "配額\n"
	command += "沒事"

	switch msg.Text {
	case "指令":
		return command
	case "帥氣指數":
		return "你今天的帥氣指數為" + fmt.Sprintf("%d", rand.Intn(100)) + "%"
	case "天氣台北":
		return weather.OpenWether(1665148, "台北")
	case "天氣台中":
		return weather.OpenWether(1668399, "台中")
	case "天氣台南":
		return weather.OpenWether(1668355, "台南")
	case "天氣中壢":
		return weather.OpenWether(1676087, "中壢")
	case "配額":
		return quotaParser(quota)
	case "沒事":
		return "沒事別吵"
	default:
		return "查無指令 <3"
	}

}
