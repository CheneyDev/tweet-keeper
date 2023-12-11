package utils

import (
	"encoding/json"
	"time"
)

// Tweet represents the structure of a tweet.
type Tweet struct {
	Likes          int
	CreatedAt      time.Time
	Status         string
	Text           string
	Retweets       int
	Bookmarks      int
	Quotes         int
	Replies        int
	Lang           string
	Views          string
	ConversationID string `json:"conversation_id"`
	Author         Author
	Media          Media
	ID             string
}

// Author represents the author of the tweet.
type Author struct {
	RestID       string `json:"rest_id"`
	Name         string
	ScreenName   string `json:"screen_name"`
	Image        string
	BlueVerified bool `json:"blue_verified"`
	SubCount     int  `json:"sub_count"`
}

// Media represents the media content in the tweet.
type Media struct {
	Photo []Photo
	Video []Video
}

// Photo represents a photo in the tweet.
type Photo struct {
	MediaURLHttps string `json:"media_url_https"`
}

// Video represents a video in the tweet.
type Video struct {
	MediaURLHttps string `json:"media_url_https"`
	Variants      []VideoVariant
}

// VideoVariant represents a variant of a video.
type VideoVariant struct {
	Bitrate     int
	ContentType string `json:"content_type"`
	URL         string
}

func HandleTweet(jsonData string) (*Tweet, error) {
	var tweet Tweet
	err := json.Unmarshal([]byte(jsonData), &tweet)
	if err != nil {
		return nil, err
	}
	return &tweet, nil
}

func GetStr() string {
	return "Heook"
}
