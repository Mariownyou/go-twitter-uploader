package main

import "github.com/mariownyou/go-twitter-uploader/twitter_uploader"

const (
	consumerKey       = "key"
	consumerSecret    = "secret"
	accessToken       = "token"
	accessTokenSecret = "token_secret"
)

func main() {
	uploader := twitter_uploader.New(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	file := []byte("file")
	uploader.Upload("text", file, "filename")
}
