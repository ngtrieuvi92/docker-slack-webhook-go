package main

import "github.com/ashwanthkumar/slack-go-webhook"
import "fmt"
import "os"

func main() {
	webhookUrl := os.Getenv("SLACK_HOOK_URL")
	if webhookUrl == "" {
		fmt.Printf("error: SLACK_HOOK_URL is quired")
		os.Exit('1')
	}
	username := os.Getenv("SLACK_USERNAME")
	if username == "" {
		username = "CI-Bot"
	}
	channel := os.Getenv("SLACK_CHANNEL")
	if channel == "" {
		channel = "#general"
	}
	icon := os.Getenv("SLACK_ICON")
	if icon == "" {
		icon = ":monkey_face:"
	}
	textMessage := os.Getenv("SLACK_TEXT_MESSAGE")
	if textMessage == "" {
		textMessage = "Build completed!"
	}

	payload := slack.Payload{
		Text:      textMessage,
		Username:  username,
		Channel:   channel,
		IconEmoji: icon,
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
