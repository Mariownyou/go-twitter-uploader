package twitter_uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/dghubble/oauth1"
)

var types = map[string]string{
	".jpg":  "tweet_image",
	".jpeg": "tweet_image",
	".png":  "tweet_image",
	".gif":  "tweet_gif",
	".mp4":  "amplify_video",
	".mov":  "amplify_video",
}

type Uploader struct {
	Client *http.Client
}

func New(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Uploader {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	return &Uploader{
		Client: httpClient,
	}
}

type Media struct {
	MediaIDs []string `json:"media_ids"`
}

type Tweet struct {
	Text  string `json:"text"`
	Media *Media `json:"media,omitempty"`
}

func (u *Uploader) Downlaod() {} // TODO

func (u *Uploader) Upload(text string, file []byte, filename string) {
	path := "https://api.twitter.com/2/tweets"

	t := types[filepath.Ext(filename)]
	if t != "twee_image" {
		return
	}

	mediaID := u.uploadImage(file, filename)

	tweet := Tweet{
		Text:  text,
		Media: &Media{MediaIDs: []string{mediaID}},
	}

	payload, _ := json.Marshal(tweet)
	reader := bytes.NewReader(payload)
	resp, _ := u.Client.Post(path, "application/json", reader)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))
}

func (u *Uploader) uploadImage(file []byte, filename string) string {
	// amplify_video, tweet_gif, tweet_image, and tweet_video
	path := "https://upload.twitter.com/1.1/media/upload.json"

	// create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// add image to multipart form
	part, _ := writer.CreateFormFile("media", filename)
	reader := bytes.NewReader(file)
	io.Copy(part, reader)
	writer.Close()
	// build request
	req, _ := http.NewRequest("POST", path, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// send request
	resp, _ := u.Client.Do(req)
	defer resp.Body.Close()
	// read response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(respBody))

	type MediaResponse struct {
		MediaID string `json:"media_id_string"`
	}

	var mediaResponse MediaResponse
	// unmarshal response
	json.Unmarshal(respBody, &mediaResponse)
	fmt.Printf("Media ID: %v\n", mediaResponse.MediaID)

	return mediaResponse.MediaID
}

func (u *Uploader) uploadVideo(filename string) {}
