package bot

import (
	"encoding/json"
	"fmt"
)

type bot struct {
	Token       string `json:"token"`
	ClientID    string `json:"client_id"`
	ClientToken string `json:"client_token"`
	Username    string `json:"username"`
	GiphyKey    string `json:"giphy_key"`
}

// NewBot creates a new instance of struct, bot
// It takes in a byte array and uses json unmarshal to set the struct
// The byte array will be read from secrets.json
func NewBot(data []byte) bot {
	var newBot bot
	err := json.Unmarshal(data, &newBot)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return newBot
}
