package commands

import "github.com/bwmarrin/discordgo"

var BasicCommand = &discordgo.ApplicationCommand{
	Name: "basic-command",
	// All commands and options must have a description
	// Commands/options without description will fail the registration
	// of the command.
	Description: "Basic command",
}

func BasicCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var userId string

	if i.User != nil {
		userId = i.User.ID
	} else {
		userId = i.Member.User.ID
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command" + userId,
		},
	})
}
