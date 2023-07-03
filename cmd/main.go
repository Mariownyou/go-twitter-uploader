package main

import (
	"os"

	"github.com/mariownyou/go-twitter-uploader/twitter_uploader"
)

const (
	consumerKey       = "key"
	consumerSecret    = "secret"
	accessToken       = "token"
	accessTokenSecret = "token_secret"
)

func main() {
	uploader := twitter_uploader.New(consumerKey, consumerSecret, accessToken, accessTokenSecret)

	file, err := os.ReadFile("video.mp4")
	if err != nil {
		panic(err)
	}

	uploader.Upload("text", file, "video.mp4")
	// uploader.UploadVideo("text", file, "video.mp4")
}
