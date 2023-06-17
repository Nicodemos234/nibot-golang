package commands

import "github.com/bwmarrin/discordgo"

// Command struct
type Command struct {
	Name string
	Exec func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	Help string
}

func GetAllCommands() map[string]Command {
	commands := make(map[string]Command)

	commands["ping"] = Command{
		Name: "Ping",
		Exec: Ping,
		Help: "Te retorna pong :)",
	}
	commands["play"] = Command{
		Name: "Play",
		Exec: Play,
		Help: "!play {link da música} para tocar uma música",
	}
	commands["rpg"] = Command{
		Name: "Rpg",
		Exec: Rpg,
		Help: "!rpg para tocar as músicas do rpg",
	}
	commands["avatar"] = Command{
		Name: "Avatar",
		Exec: Avatar,
		Help: "Retorna o avatar de alguém ou o seu",
	}
	commands["guilds"] = Command{
		Name: "Guilds",
		Exec: Guilds,
		Help: "Lista todos os servidores que o bot está",
	}

	return commands
}
