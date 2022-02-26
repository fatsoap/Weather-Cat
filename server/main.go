package main

import (
	"ez_line_bot/robbot"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file")
		return
	}
	srv := robbot.Init()
	PORT := ":" + os.Getenv("PORT")
	http.ListenAndServe(PORT, srv)
}
