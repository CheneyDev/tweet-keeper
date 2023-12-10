package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"twitter-to-telegram/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	rapidApiKey := os.Getenv("RAPID_API_KEY")
	rapidApiHost := os.Getenv("RAPID_API_HOST")

	tweetUrl := ""

	res, _ := api.GetTweet(tweetUrl, rapidApiKey, rapidApiHost)

	//print
	fmt.Println(res)
}
