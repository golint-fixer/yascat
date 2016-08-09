package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	// "strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Token        string
	Client_ID    string
	Client_Token string
	Username     string
}

var BotID string
var buffer = make([][]byte, 0)

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

	err := loadSound()
	if err != nil {
		fmt.Println("Error loading sound:", err)
		return
	}


	// Handlers
	dg.AddHandler(yasSay)

	// Open discord connection and run bot
	e = dg.Open()
	if e != nil {
		fmt.Println("Error opening connection", e)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}


func loadSound() error {
	file, e := os.Open("yascat.dca")

	if e != nil {
		fmt.Println("Error opening dca file:", e)
		return e
	}

	var opuslen int16

	for {
		// Get opus frame lenght from dca file
		e = binary.Read(file, binary.LittleEndian, &opuslen)

		if e == io.EOF || e == io.ErrUnexpectedEOF {
			return nil
		}

		if e != nil {
			fmt.Println("Error reading from dca file:", e)
			return e
		}

		InBuf := make([]byte, opuslen)
		e = binary.Read(file, binary.LittleEndian, &InBuf)

		if e != nil {
			fmt.Println("Error reading from dca file:", e)
			return e
		}

		buffer = append(buffer, InBuf)
	}
}


func playSound(s *discordgo.Session, guildID, channelID string) (err error) {

	vc, e := s.ChannelVoiceJoin(guildID, channelID, false, true)
		if e != nil {
			return
		}

		_ = vc.Speaking(true)

		for _, buff := range buffer {
			vc.OpusSend <- buff
		}

		_ = vc.Speaking(false)

		time.Sleep(30 * time.Millisecond)

		_ = vc.Disconnect()

		return
}


// Message event handle
func yasSay(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "!yas" {
		// _, _ = s.ChannelMessageSend(m.ChannelID, "yas")
		c, e := s.State.Channel(m.ChannelID)
		if e != nil {
			fmt.Println("Could not find channel:", e)
			return
		}


		g, e := s.State.Guild(c.GuildID)
		if e != nil {
			fmt.Println("Could not find guild", e)
			return
		}


		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				e := playSound(s, g.ID, vs.ChannelID)
				if e != nil {
					fmt.Println("Error playing sound", e)
				}
				return
			}
		}
	}
}
