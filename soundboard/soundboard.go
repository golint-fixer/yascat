package soundboard

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Sound commands mapped to sound file
var Commands = map[string]string{
	"!yas":      "assets/yascat.dca",
	"!bonglord": "assets/bonglord.dca",
}

// Sound struct that provides the name of the sound and a buffer
type Sound struct {
	Name   string
	buffer [][]byte
}

// Given a fileName, return the Sound struct with loaded buffer for sound
func LoadSound(fileName string) (s Sound, err error) {
	// Instantiate a sound with a new buffer and name
	newSound := Sound{
		fileName,
		make([][]byte, 0),
	}

	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("Error opening dca file:", e)
		return newSound, e
	}

	var opuslen int16
	for {
		// Get opus frame length from dca file
		e = binary.Read(file, binary.LittleEndian, &opuslen)
		if e != nil {
			fmt.Println("Error reading opus frame length from dca file:", e)
			return newSound, e
		}

		// Create a buffer and read PCM packets into it
		InBuf := make([]byte, opuslen)
		e = binary.Read(file, binary.LittleEndian, &InBuf)
		if e != nil {
			fmt.Println("Error reading from dca file:", e)
			return newSound, e
		}

		// Append the sound to the buffer in newSound
		newSound.buffer = append(newSound.buffer, InBuf)
	}
	return newSound, nil
}

func PlaySound(s *discordgo.Session, guildID, channelID string, sound Sound) (err error) {

	vc, e := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if e != nil {
		return e
	}

	_ = vc.Speaking(true)

	for _, buff := range sound.buffer {
		vc.OpusSend <- buff
	}
	_ = vc.Speaking(false)
	time.Sleep(30 * time.Millisecond)
	_ = vc.Disconnect()
	return
}
