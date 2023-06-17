package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Avatar(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	resolution := "2048"
	avatar := m.Author.AvatarURL(resolution)
	if len(args) > 0 {
		mentionedUser := strings.TrimSuffix(strings.TrimPrefix(args[0], "<@"), ">")
		user, err := s.User(mentionedUser)
		if err == nil {
			avatar = user.AvatarURL(resolution)
		}
	}
	s.ChannelMessageSend(m.ChannelID, avatar)
}
