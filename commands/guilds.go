package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Guilds(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	guildNames := []string{}
	for _, guild := range s.State.Guilds {
		guildNames = append(guildNames, guild.Name)
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, strings.Join(guildNames, "\n"))
}
