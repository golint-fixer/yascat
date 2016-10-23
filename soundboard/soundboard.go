package soundboard

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// Sound commands mapped to sound file
var Commands = map[string]string{
	"!yas":      "assets/yascat.dca",
	"!bonglord": "assets/bonglord.dca",
}

type Sound struct {
	Name   string
	buffer [][]byte
}

// Given a fileName, return the Sound struct with loaded buffer for sound
func loadSound(fileName string) (s *Sound, err error) {
	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("Error opening dca file:", e)
		return nil, e
	}

	// Instantiate a sound with a new buffer and name
	newSound := Sound{
		command,
		make([][]byte, 0),
	}

	var opuslen int16
	for {
		// Get opus frame length from dca file
		e = binary.Read(file, binary.LittleEndian, &opuslen)
		if e != nil {
			fmt.Println("Error reading from dca file:", e)
			return nil, e
		}

		// Create a buffer and read PCM packets into it
		InBuf := make([]byte, opuslen)
		e = binary.Read(file, binary.LittleEndian, &InBuf)
		if e != nil {
			fmt.Println("Error reading from dca file:", e)
			return nil, e
		}

		// Append the sound to the buffer in newSound
		newSound.buffer = append(newSound.buffer, InBuf...)
	}
	return newSound, nil
}

func playSound(s *discordgo.Session, guildID, channelID string, sound Sound) (err error) {

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
