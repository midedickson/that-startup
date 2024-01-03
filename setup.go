package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

var requester *http.Client

func createHttpClientFromOAuth() {
	config := oauth1.NewConfig(os.Getenv("TWITTER_API"), os.Getenv("TWITTER_API_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	requester = config.Client(context.Background(), token)
}

func setUp() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	createHttpClientFromOAuth()
}
