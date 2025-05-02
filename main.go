package main

import (
	"log"
	"net/http"
	"os"
	"time"

	bot "atreids.com/go-spin-the-wheel/bot"
	"github.com/joho/godotenv"
)

func checkNetwork() bool {
	log.Print("Checking network connection.")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if !checkNetwork() {
		log.Fatal("Network connection check failed. Exiting...")
	}

	botToken := os.Getenv("BOT_TOKEN")

	bot.BotToken = botToken
	bot.Run()
}
