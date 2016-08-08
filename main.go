package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	// "strings"
	// "time"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Token        string
	Client_ID    string
	Client_Token string
	Username     string
}

var BotID string

func main() {

	// Retrieve bot information from secrets
	sfile, e := os.Open("secrets.json")
	if e != nil {
		fmt.Println("Error opening secrets file:", e)
		return
	}
	defer sfile.Close()
	sdata, e := ioutil.ReadAll(sfile)
	if e != nil {
		fmt.Println("Error reading secrets file:", e)
		return
	}

	var bot Bot
	json.Unmarshal(sdata, &bot)

	// Extract the token from the bot and create new discord session
	var token string = bot.Token
	dg, e := discordgo.New(token)
	if e != nil {
		fmt.Println("Error creating Discord session, e")
		return
	}

	u, e := dg.User("@me")
	if e != nil {
		fmt.Println("Error obtaining account details,", e)
	}

	BotID = u.ID

	dg.AddHandler(yasSay)

	e = dg.Open()
	if e != nil {
		fmt.Println("Error opening connection", e)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}

// Message event handle
func yasSay(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "!yas" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "yas")
	}
}
