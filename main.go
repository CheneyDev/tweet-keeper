package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
	"log"
	"os"
	"time"
)

type RequestBody struct {
	TweetURL string `json:"tweetUrl"`
	AuthKey  string `json:"authKey"`
}

func main() {
	// 加载环境变量，如果不存在则使用默认值
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file, using default values")
	}
	notionToken := os.Getenv("NOTION_TOKEN")
	pageID := os.Getenv("NOTION_PAGE_ID")

	if notionToken == "" || pageID == "" {
		log.Fatal("NOTION_TOKEN and NOTION_PAGE_ID must be set")
	}
	notionClient := notionapi.NewClient(notionapi.Token(notionToken))

	// 创建一个新的文本段落块
	// Assume we have the parent block ID to which we want to append a child block.
	parentBlockID := notionapi.BlockID(pageID)

	currentTime := time.Now()

	// Define the paragraph block we want to add.
	paragraphBlock := notionapi.ParagraphBlock{
		Paragraph: notionapi.Paragraph{
			RichText: []notionapi.RichText{
				{
					Type: notionapi.ObjectTypeText,
					Text: &notionapi.Text{
						Content: "Hello from Notion APIssss!",
					},
				},
			},
		},
	}
	imageBlock := notionapi.ImageBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:           "image",
			Object:         "block",
			CreatedTime:    &currentTime,
			LastEditedTime: &currentTime,
		},
		Image: notionapi.Image{
			Type: "external",
			External: &notionapi.FileObject{
				URL: "https://f005.backblazeb2.com/file/nsfw-twitter/IMG_9961.JPG",
			},
		},
	}
	// Set the type of the block.
	paragraphBlock.Type = "paragraph"

	// Set the base block properties.
	paragraphBlock.Object = "block"
	paragraphBlock.HasChildren = false
	paragraphBlock.CreatedTime = &currentTime
	paragraphBlock.LastEditedTime = &currentTime

	// Append the new block as a child to the parent block.
	appendBlockRequest := &notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{&paragraphBlock},
	}

	appendImageBlockRequest := &notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{&imageBlock},
	}

	response, err := notionClient.Block.AppendChildren(context.Background(), parentBlockID, appendBlockRequest)
	if err != nil {
		log.Fatalf("Failed to append block: %v\n", err)
	}

	response, err = notionClient.Block.AppendChildren(context.Background(), parentBlockID, appendImageBlockRequest)
	if err != nil {
		log.Fatalf("Failed to append image block: %v\n", err)
	}
	log.Printf("Appended block with response: %+v\n", response)

}

//func main() {
//	// 加载环境变量，如果不存在则使用默认值
//	if err := godotenv.Load(); err != nil {
//		log.Println("Warning: Error loading .env file, using default values")
//	}
//
//	// 获取S3桶名，如果环境变量不存在则使用默认桶名
//	bucketName := os.Getenv("S3_BUCKET_NAME")
//	if bucketName == "" {
//		log.Fatal("S3_BUCKET_NAME environment variable is not set.")
//	}
//
//	// 创建S3客户端实例
//	s3Client := s3_storage.NewS3Client()
//
//	// 打开文件
//	file, err := os.Open("test.txt")
//	if err != nil {
//		log.Fatalf("Error opening file: %s", err)
//	}
//	defer file.Close() // 确保文件最终会被关闭
//
//	// 上传文件到S3
//	_, err = s3Client.Service.PutObject(&s3.PutObjectInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String("./tt/testhhh.txt"),
//		Body:   file,
//	})
//	if err != nil {
//		log.Fatalf("Failed to upload object to S3: %s", err)
//	}
//
//	// 列出S3桶中的对象
//	input := &s3.ListObjectsV2Input{
//		Bucket: aws.String(bucketName),
//	}
//	result, err := s3Client.Service.ListObjectsV2(input)
//	if err != nil {
//		log.Fatalf("Unable to list objects in bucket: %s", err)
//	}
//
//	// 打印结果和桶名
//	fmt.Println("List of objects in bucket:", result)
//}

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
