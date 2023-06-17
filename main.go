package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nibot-discord/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string
	BotPrefix string
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

// Command struct
type Command struct {
	Name string
	Exec func()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if strings.HasPrefix(m.Content, BotPrefix) {
		withoutPrefix := strings.TrimPrefix(m.Content, BotPrefix)
		split := strings.Split(withoutPrefix, " ")
		commandName := split[0]
		args := split[1:]

		fmt.Println("First value:", commandName)
		fmt.Println("Rest of the values:", args)

		commands := commands.GetAllCommands()
		cmd, ok := commands[commandName]
		if ok {
			cmd.Exec(s, m, args)
		}
	}
}

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
	return
}
