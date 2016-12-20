package gif_search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	// "path/filepath"
	"strings"
	"time"

	"yascat/bot"
)

type SearchData struct {
	Data       []ImageData `json:"data"`
	Meta       Meta        `json:"meta"`
	Pagination Pagination  `json:"pagination"`
}

type ImageData struct {
	Img_type          string                 `json:"type"`
	Id                string                 `json:"id"`
	Slug              string                 `json:"slug"`
	Url               string                 `json:"url"`
	Bitly_gif_url     string                 `json:"bitly_gif_url"`
	Bitly_url         string                 `json:"bitly_url"`
	Embed_url         string                 `json:"embed_url"`
	Username          string                 `json:"username"`
	Source            string                 `json:"source"`
	Rating            string                 `json:"rating"`
	Caption           string                 `json:"caption"`
	Content_url       string                 `json:"content_url"`
	Source_tld        string                 `json:"source_tld"`
	Source_post_url   string                 `json:"source_post_url"`
	Import_datetime   string                 `json:"import_datetime"`
	Trending_datetime string                 `json:"trending_datetime"`
	Images            map[string]interface{} `json:"images"`
}

type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type Pagination struct {
	Total_count int `json:"total_count"`
	Count       int `json:"count"`
	Offset      int `json:"offset"`
}

// Make a request to giphy given a list of search terms
func GetGif(search_terms []string) (gif string) {
	// TODO(doria): This reuses the bot extraction... need to refactor
	sfile, e := os.Open("secrets.json")
	if e != nil {
		fmt.Println("Error opening secrets file:", e)
	}
	defer sfile.Close()
	sdata, e := ioutil.ReadAll(sfile)
	if e != nil {
		fmt.Println("Error reading secrets file:", e)
		return
	}

	bot := bot.NewBot(sdata)

	query := strings.Join(search_terms, "+")
	queryUrl := "http://api.giphy.com/v1/gifs/search?q=" + query + "&api_key=" + bot.GiphyKey
	resp, e := http.Get(queryUrl)
	if e != nil {
		fmt.Println("Unable to get gif")
	}
	defer resp.Body.Close()
	searchResults := SearchData{}
	e = json.NewDecoder(resp.Body).Decode(&searchResults)
	if e != nil {
		fmt.Println("Unable to unmarshal json")
	}

	if len(searchResults.Data) < 1 {
		return "Meow, cannot find any gifs."
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(searchResults.Data)
	return searchResults.Data[n].Url
}
