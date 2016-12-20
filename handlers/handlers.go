package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"yascat/gif_search"
	"yascat/soundboard"
)

// Discord session handler for yascat soundboard
func SoundboardHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO(doria): Add logging
	if fileName, ok := soundboard.Commands[m.Content]; ok {
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
				sound, e := soundboard.LoadSound(fileName)
				e = soundboard.PlaySound(s, g.ID, vs.ChannelID, sound)
				if e != nil {
					fmt.Println("Error playing sound", e)
				}
				return
			}
		}
	}
}

// Discord session handler for yascat gif_search
func GifSearchHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO(doria): Add logging
	message := strings.Split(m.ContentWithMentionsReplaced(), " ")
	fmt.Println(message)
	fmt.Println(len(message))
	if len(message) > 2 && message[0] == "@yascat" && message[1] == "!findgif" {
		gif_url := gif_search.GetGif(message[2:])
		c, e := s.ChannelMessageSend(m.ChannelID, gif_url)
		if e != nil {
			fmt.Println("Could not find channel:", e)
			return
		}
		fmt.Println(c)
	}
	return
}
