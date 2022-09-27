package commands

import "github.com/bwmarrin/discordgo"

var Subcommands = &discordgo.ApplicationCommand{
	Name:        "subcommands",
	Description: "Subcommands and command groups example",
	Options: []*discordgo.ApplicationCommandOption{
		// When a command has subcommands/subcommand groups
		// It must not have top-level options, they aren't accesible in the UI
		// in this case (at least not yet), so if a command has
		// subcommands/subcommand any groups registering top-level options
		// will cause the registration of the command to fail

		{
			Name:        "subcommand-group",
			Description: "Subcommands group",
			Options: []*discordgo.ApplicationCommandOption{
				// Also, subcommand groups aren't capable of
				// containing options, by the name of them, you can see
				// they can only contain subcommands
				{
					Name:        "nested-subcommand",
					Description: "Nested subcommand",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
			Type: discordgo.ApplicationCommandOptionSubCommandGroup,
		},
		// Also, you can create both subcommand groups and subcommands
		// in the command at the same time. But, there's some limits to
		// nesting, count of subcommands (top level and nested) and options.
		// Read the intro of slash-commands docs on Discord dev portal
		// to get more information
		{
			Name:        "subcommand",
			Description: "Top-level subcommand",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func SubcommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	content := ""

	// As you can see, names of subcommands (nested, top-level)
	// and subcommand groups are provided through the arguments.
	switch options[0].Name {
	case "subcommand":
		content = "The top-level subcommand is executed. Now try to execute the nested one."
	case "subcommand-group":
		options = options[0].Options
		switch options[0].Name {
		case "nested-subcommand":
			content = "Nice, now you know how to execute nested commands too"
		default:
			content = "Oops, something went wrong.\n" +
				"Hol' up, you aren't supposed to see this message."
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
