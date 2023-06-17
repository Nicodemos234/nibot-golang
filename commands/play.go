package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func Play(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "Play executing")

	voiceChannel := ""
	// TODO: Check a better way todo this search
	for _, guild := range s.State.Guilds {
		if guild.ID == m.GuildID {
			voiceStates := guild.VoiceStates
			voiceChannel = voiceStates[0].ChannelID
		}
	}
	if voiceChannel == "" {
		return
	}
	dgv, err := s.ChannelVoiceJoin(m.GuildID, voiceChannel, false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start loop and attempt to play all files in the given folder
	folder := "./audios/rpg"
	// fmt.Println("Reading Folder: ", folder)
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {
		fmt.Println("PlayAudioFile:", f.Name())
		dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", folder, f.Name()), make(chan bool))
	}

}
