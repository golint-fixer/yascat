package gif_search

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	query := strings.Join(search_terms, "+")
	queryUrl := "http://api.giphy.com/v1/gifs/search?q=" + query + "&api_key=dc6zaTOxFJmzC"
	resp, err := http.Get(queryUrl)
	if err != nil {
		fmt.Println("Unable to get gif")
	}
	defer resp.Body.Close()
	searchResults := SearchData{}
	err = json.NewDecoder(resp.Body).Decode(&searchResults)
	if err != nil {
		fmt.Println("Unable to unmarshal json")
	}
	fmt.Println(searchResults)

	if len(searchResults.Data) < 1 {
		return "Meow, cannot find any gifs."
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(searchResults.Data)
	return searchResults.Data[n].Url
}
