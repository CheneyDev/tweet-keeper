package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"log"
	"os"
	"tweet-keeper/service/s3-storage"
)

type RequestBody struct {
	TweetURL string `json:"tweetUrl"`
	AuthKey  string `json:"authKey"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	bucketName := os.Getenv("S3_BUCKET_NAME")
	s3Client := s3_storage.NewS3Client()
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	result, err := s3Client.Service.ListObjectsV2(input)
	if err != nil {
		log.Fatalf("无法列出桶中的对象: %s", err)
	}
	//print
	fmt.Println(result, bucketName)

}

//func main() {
//	rapidApiKey := os.Getenv("RAPID_API_KEY")
//	rapidApiHost := os.Getenv("RAPID_API_HOST")
//	expectedAuthKey := os.Getenv("AUTH_KEY")
//
//	handler := func(w http.ResponseWriter, r *http.Request) {
//		var reqBody RequestBody
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			http.Error(w, "Error reading request body", http.StatusInternalServerError)
//			return
//		}
//		defer r.Body.Close()
//
//		err = json.Unmarshal(body, &reqBody)
//		if err != nil {
//			http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
//			return
//		}
//
//		if reqBody.TweetURL == "" {
//			http.Error(w, "tweetUrl is required", http.StatusBadRequest)
//			return
//		}
//
//		if expectedAuthKey != reqBody.AuthKey {
//			http.Error(w, "Invalid authentication key", http.StatusUnauthorized)
//			return
//		}
//
//		res, err := api.GetTweet(reqBody.TweetURL, rapidApiKey, rapidApiHost)
//		if err != nil {
//			http.Error(w, fmt.Sprintf("Error getting tweet: %v", err), http.StatusInternalServerError)
//			return
//		}
//		tweet, err := utils.HandleTweet(res)
//		tweetText := fmt.Sprintf("%s\n%s", tweet.Text, tweet.Author.ScreenName)
//		tweetPhotos := tweet.Media.Photo
//		tweetVideos := tweet.Media.Video
//		// 打印结果
//		fmt.Fprintf(w, "Response: %v\n%v\n%v\n", tweetText, tweetPhotos, tweetVideos)
//	}
//	http.HandleFunc("/getTweet", handler)
//
//	fmt.Println("Server is starting...")
//	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
//}
