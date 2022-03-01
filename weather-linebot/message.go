package main

import (
	"fmt"
	"time"

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

func handleMessage(bot *linebot.Client, event *linebot.Event) error {
	text := ""
	altText := "貓咪"
	switch message := event.Message.(type) {
	case *linebot.LocationMessage: // 位置訊息
		lat := message.Latitude
		lon := message.Longitude
		fmt.Printf("User:%s\nLat:%.0f\nLon:%.0f", event.Source.UserID, lat, lon)
		// Open Weather Map api
		current, err := owm.NewCurrent("C", "zh_tw", env.OWM_API_KEY)
		if err != nil {
			fmt.Printf("Init Open Weather Map Failed")
		}
		cords := owm.Coordinates{Latitude: lat, Longitude: lon}
		current.CurrentByCoordinates(&cords)
		{
			text = fmt.Sprintf(flexMessageTemplate,
				fmt.Sprintf("%v", time.Now().Nanosecond())[:2],
				current.Weather[0].Description,
				current.Main.Temp,
				current.Main.TempMax,
				current.Main.TempMin,
				handleWind(current.Wind.Speed, current.Wind.Deg),
				current.Clouds.All,
				current.Main.Humidity,
				handleTime(current.Sys.Sunrise),
				handleTime(current.Sys.Sunset),
				current.Main.FeelsLike,
				current.Main.Pressure,
				current.Rain.OneH,
				current.Snow.OneH,
			)
		}
		altText = "天氣"
	default: // 其他訊息
		fmt.Println("Not Location Message")
		text = fmt.Sprintf(hintFlexMessageTemplate, fmt.Sprintf("%v", time.Now().Nanosecond())[:2])
	}
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(text))
	if err != nil {
		fmt.Println("Flex Message Unmarshal Failed")
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage(altText, flexContainer)).Do(); err != nil {
		fmt.Println("Reply Flex Message Failed")
	}
	return nil
}

func handleWind(speed float64, deg float64) string {
	dir := ""
	{
		if 384.76 <= deg || deg <= 11.25 {
			dir = "北"
		} else if 11.26 <= deg && deg <= 33.75 {
			dir = "北東北"
		} else if 33.76 <= deg && deg <= 56.25 {
			dir = "東北"
		} else if 56.26 <= deg && deg <= 78.75 {
			dir = "東東北"
		} else if 78.76 <= deg && deg <= 101.25 {
			dir = "東"
		} else if 101.26 <= deg && deg <= 123.75 {
			dir = "東東南"
		} else if 123.76 <= deg && deg <= 146.25 {
			dir = "東南"
		} else if 146.26 <= deg && deg <= 168.75 {
			dir = "南東南"
		} else if 168.76 <= deg && deg <= 191.25 {
			dir = "南"
		} else if 191.26 <= deg && deg <= 213.75 {
			dir = "南西南"
		} else if 213.76 <= deg && deg <= 236.25 {
			dir = "西南"
		} else if 236.26 <= deg && deg <= 258.75 {
			dir = "西西南"
		} else if 258.76 <= deg && deg <= 281.25 {
			dir = "西"
		} else if 281.26 <= deg && deg <= 303.75 {
			dir = "西西北"
		} else if 303.76 <= deg && deg <= 326.25 {
			dir = "西北"
		} else if 326.26 <= deg && deg <= 348.75 {
			dir = "北西北"
		}
	}
	dir += fmt.Sprintf(" %v 公尺/秒", speed)
	if speed <= 0.2 {
		dir = "靜風"
	}
	return dir
}

func handleTime(timestamp int) string {
	tz := time.FixedZone("UTC+8", 8*60*60)
	ts := time.Unix(int64(timestamp), 0)
	return fmt.Sprintf("%02d:%02d", ts.In(tz).Hour(), ts.In(tz).Minute())
}

var flexMessageTemplate = `{
	"type": "bubble",
	"hero": {
	  "type": "image",
	  "url": "https://source.unsplash.com/random/?weather,%v",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"uri": "http://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [],
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "%v",
				  "size": "md",
				  "align": "center",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "%.0f˚",
				  "size": "xxl",
				  "align": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [],
				  "flex": 1
				}
			  ],
			  "flex": 0,
			  "alignItems": "center",
			  "justifyContent": "center"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "margin": "none",
			  "spacing": "none",
			  "contents": [
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [],
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "最高",
					  "color": "#aaaaaa",
					  "size": "sm",
					  "flex": 1
					},
					{
					  "type": "text",
					  "text": "%.0f˚",
					  "wrap": true,
					  "color": "#666666",
					  "size": "sm",
					  "flex": 1
					}
				  ],
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "最低",
					  "color": "#aaaaaa",
					  "size": "sm",
					  "flex": 1
					},
					{
					  "type": "text",
					  "text": "%0.f˚",
					  "wrap": true,
					  "color": "#666666",
					  "size": "sm",
					  "flex": 1
					}
				  ],
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [],
				  "flex": 1
				}
			  ]
			}
		  ]
		},
		{
		  "type": "box",
		  "layout": "horizontal",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "風",
			  "color": "#aaaaaa",
			  "size": "sm",
			  "flex": 1
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "none",
			  "contents": [
				{
				  "type": "text",
				  "text": "%s",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "center"
				}
			  ],
			  "flex": 3
			}
		  ],
		  "alignItems": "center",
		  "justifyContent": "center"
		},
		{
		  "type": "box",
		  "layout": "horizontal",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "雲量",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "%v%%",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end",
				  "offsetEnd": "sm"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "濕度",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1,
				  "offsetStart": "sm"
				},
				{
				  "type": "text",
				  "text": "%v%%",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end"
				}
			  ]
			}
		  ],
		  "alignItems": "center",
		  "justifyContent": "center"
		},
		{
		  "type": "box",
		  "layout": "horizontal",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "日出",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "%s",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end",
				  "offsetEnd": "sm"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "日落",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1,
				  "offsetStart": "sm"
				},
				{
				  "type": "text",
				  "text": "%s",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end"
				}
			  ]
			}
		  ],
		  "alignItems": "center",
		  "justifyContent": "center"
		},
		{
		  "type": "box",
		  "layout": "horizontal",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "體感",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "%.0f˚",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end",
				  "offsetEnd": "sm"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "氣壓",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1,
				  "offsetStart": "sm"
				},
				{
				  "type": "text",
				  "text": "%0.f百帕",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end"
				}
			  ]
			}
		  ],
		  "alignItems": "center",
		  "justifyContent": "center"
		},
		{
		  "type": "box",
		  "layout": "horizontal",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "降雨量",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "%.0fmm",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end",
				  "offsetEnd": "sm"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "降雪量",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1,
				  "offsetStart": "sm"
				},
				{
				  "type": "text",
				  "text": "%.0fmm",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 1,
				  "align": "end"
				}
			  ]
			}
		  ],
		  "alignItems": "center",
		  "justifyContent": "center"
		}
	  ]
	}
  }`

var hintFlexMessageTemplate = `{
	"type": "bubble",
	"hero": {
	  "type": "image",
	  "url": "https://source.unsplash.com/random/?cat,%v",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"uri": "http://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "md",
	  "contents": [
		{
		  "type": "text",
		  "text": "使用位置訊息來取得當前位置的天氣",
		  "wrap": true,
		  "weight": "regular",
		  "gravity": "center",
		  "size": "sm",
		  "align": "center"
		},
		{
		  "type": "text",
		  "text": "Send location message to get current weather (=ↀωↀ=)",
		  "wrap": true,
		  "weight": "regular",
		  "gravity": "center",
		  "size": "sm",
		  "align": "center"
		},
		{
		  "type": "button",
		  "action": {
			"type": "uri",
			"label": "Github",
			"uri": "https://github.com/fatsoap/Weather-Cat"
		  }
		}
	  ]
	}
  }`
