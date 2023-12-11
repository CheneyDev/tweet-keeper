package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"twitter-to-telegram/api"
	"twitter-to-telegram/utils"
)

type RequestBody struct {
	TweetURL string `json:"tweetUrl"`
	AuthKey  string `json:"authKey"`
}

func main() {
	rapidApiKey := os.Getenv("RAPID_API_KEY")
	rapidApiHost := os.Getenv("RAPID_API_HOST")
	expectedAuthKey := os.Getenv("AUTH_KEY")

	handler := func(w http.ResponseWriter, r *http.Request) {
		var reqBody RequestBody
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
			return
		}

		if reqBody.TweetURL == "" {
			http.Error(w, "tweetUrl is required", http.StatusBadRequest)
			return
		}

		if expectedAuthKey != reqBody.AuthKey {
			http.Error(w, "Invalid authentication key", http.StatusUnauthorized)
			return
		}

		res, err := api.GetTweet(reqBody.TweetURL, rapidApiKey, rapidApiHost)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting tweet: %v", err), http.StatusInternalServerError)
			return
		}
		tweet, err := utils.HandleTweet(res)
		tweetText := fmt.Sprintf("%s\n%s", tweet.Text, tweet.Author.ScreenName)
		tweetPhotos := tweet.Media.Photo
		tweetVideos := tweet.Media.Video
		// 打印结果
		fmt.Fprintf(w, "Response: %v\n%v\n%v\n", tweetText, tweetPhotos, tweetVideos)
	}
	http.HandleFunc("/getTweet", handler)

	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
