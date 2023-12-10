package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTweet(tweetUrl, apiKey, rapidApiHost) (string, error) {
	tweetID, err := extractTweetID(tweetUrl)
	url := fmt.Sprintf("https://twitter-api45.p.rapidapi.com/tweet.php?id=%s", tweetId)

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add required headers
	req.Header.Add("X-Rapidapi-Key", apiKey)
	req.Header.Add("X-Rapidapi-Host", rapidApiHost)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Decode JSON
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Convert the decoded result back to JSON string for pretty print
	jsonString, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

func extractTweetID(url string) (string, error) {
    re := regexp.MustCompile(`status/(\d+)`)
    matches := re.FindStringSubmatch(url)
    if len(matches) < 2 {
        return "", fmt.Errorf("no tweet ID found in URL")
    }
    return matches[1], nil
}
