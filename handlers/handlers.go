package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"yascat/soundboard"
)

// Discord session handler for yascat soundboard
func SoundboardHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO(doria): Add logging
	if _, ok := soundboard.Commands[m.Content]; ok {
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
				sound, e := soundboard.loadSound(m.Content)
				e := soundboard.playSound(s, g.ID, vs.ChannelID, sound)
				if e != nil {
					fmt.Println("Error playing sound", e)
				}
				return
			}
		}
	}
}
