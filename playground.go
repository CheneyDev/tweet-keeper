package main

import (
    "fmt"
    "regexp"
)

func main() {
    url := "https://x.com/seitenkanna/status/1733712001212572158?s=10&t=63gEP94KrmPWz6DrhNfD1g"
    tweetID, err := extractTweetID(url)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Tweet ID:", tweetID)
    }
}

func extractTweetID(url string) (string, error) {
    re := regexp.MustCompile(`status/(\d+)`)
    matches := re.FindStringSubmatch(url)
    if len(matches) < 2 {
        return "", fmt.Errorf("no tweet ID found in URL")
    }
    return matches[1], nil
}
