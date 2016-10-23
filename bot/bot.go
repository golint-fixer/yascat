package bot

import (
	"encoding/json"
	"fmt"
)

type bot struct {
	Token        string
	Client_ID    string
	Client_Token string
	Username     string
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
