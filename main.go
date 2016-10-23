package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"

	"yascat/bot"
	"yascat/handlers"
)

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

	// Create new bot
	bot := bot.NewBot(sdata)

	// Extract the token from the bot and create new discord session
	token := bot.Token
	dg, e := discordgo.New(token)
	if e != nil {
		fmt.Println("Error creating Discord session", e)
		return
	}

	// TODO(doria): Add logging here
	// Extract the account details for bot
	// u, e := dg.User("@me")
	// if e != nil {
	// fmt.Println("Error obtaining account details,", e)
	// }

	// Add all bot handlers
	dg.AddHandler(handlers.SoundboardHandler)

	// Open discord connection and run bot
	e = dg.Open()
	if e != nil {
		fmt.Println("Error opening connection", e)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	// TODO(doria): Add logging here
	return
}
