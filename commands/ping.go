package commands

import "github.com/bwmarrin/discordgo"

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
}
