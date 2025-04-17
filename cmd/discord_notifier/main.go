package main

import (
	"flag"
	"fmt"
	"github.com/eyprz/discord_notifier/internal/config"
	"github.com/eyprz/discord_notifier/internal/webhook"
)

func main() {
	config, err := config.LoadConfig("config/config.yaml")

	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var (
		url      = flag.String("w", "", "the webhook url")
		avatar   = flag.String("a", "", "the avatar url")
		username = flag.String("u", "", "the username of the message")
	)
	flag.Parse()
	args := flag.Args()
	if flag.NArg() != 1 {
		fmt.Println("Usage: notify [-w <webhook url>] [-a <avatar url>] [-u <username>] <message>")
		fmt.Println("If you don't provide a option, the config.yaml will be used")
		return
	}
	if *url == "" {
		*url = config.WebhookURL
	}
	if *avatar == "" {
		*avatar = config.AvatarURL
	}
	if *username == "" {
		*username = config.Username
	}
	message := args[0]

	err = webhook.SendMessageToDiscord(*url, message, *avatar, *username)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}
