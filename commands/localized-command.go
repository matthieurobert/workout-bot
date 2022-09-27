package commands

import "github.com/bwmarrin/discordgo"

var LocalizedCommand = &discordgo.ApplicationCommand{
	Name:        "localized-command",
	Description: "Localized command. Description and name may vary depending on the Language setting",
	NameLocalizations: &map[discordgo.Locale]string{
		discordgo.ChineseCN: "本地化的命令",
	},
	DescriptionLocalizations: &map[discordgo.Locale]string{
		discordgo.ChineseCN: "这是一个本地化的命令",
	},
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "localized-option",
			Description: "Localized option. Description and name may vary depending on the Language setting",
			NameLocalizations: map[discordgo.Locale]string{
				discordgo.ChineseCN: "一个本地化的选项",
			},
			DescriptionLocalizations: map[discordgo.Locale]string{
				discordgo.ChineseCN: "这是一个本地化的选项",
			},
			Type: discordgo.ApplicationCommandOptionInteger,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name: "First",
					NameLocalizations: map[discordgo.Locale]string{
						discordgo.ChineseCN: "一的",
					},
					Value: 1,
				},
				{
					Name: "Second",
					NameLocalizations: map[discordgo.Locale]string{
						discordgo.ChineseCN: "二的",
					},
					Value: 2,
				},
			},
		},
	},
}

func LocalizedCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	responses := map[discordgo.Locale]string{
		discordgo.ChineseCN: "你好！ 这是一个本地化的命令",
	}
	response := "Hi! This is a localized message"
	if r, ok := responses[i.Locale]; ok {
		response = r
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
	if err != nil {
		panic(err)
	}
}
