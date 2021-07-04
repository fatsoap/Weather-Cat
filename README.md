# Weather Line Bot with Go :zap:

## Demo :rocket:

![demo](demo.gif)

## Using :hammer:

### ![](https://img.shields.io/badge/-Golang-00ADD8?logo=go&logoColor=white) :mouse:

> Golang as backend server 

### Line Bot :construction_worker:

> [Go line-bot SDK](https://github.com/line/line-bot-sdk-go)

> [Line Message API Doc](https://developers.line.biz/en/reference/messaging-api/)

### Open Weather Map API :cloud:

> [OpenWeather](https://openweathermap.org/)

> [Go API](https://github.com/briandowns/openweathermap)

## Install all package

```
go get -u ./...
```

## Add .env file

```
PORT
LINE_ACCESS_TOKEN
LINE_SCRET_TOLEN
OPENWEATHER
```

## Run

```
go run main.go
```