package commands

import (
	"encoding/json"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/matthieurobert/workout-bot/models"
)

var integerMinLenValue = 0

var Time = &discordgo.ApplicationCommand{
	Name:        "time",
	Description: "Add an exercice based on time",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Description: "Name of the exercice",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        "length",
			Description: "Duration of the exercice",
			Required:    true,
			// MinValue:    &integerMinLenValue,
		},
	},
}

func TimeHandlers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var name string
	var length int

	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if opt, ok := optionMap["name"]; ok {
		name = opt.StringValue()
	}

	if opt, ok := optionMap["length"]; ok {
		length = int(opt.IntValue())
	}

	exercice := models.Exercice{
		Id:      uuid.New(),
		Name:    name,
		Type:    "time",
		Success: true,
		Properties: models.ExerciceProperties{
			Length: length,
		},
		CreatedAt: time.Now(),
	}

	exerciceJson, err := json.Marshal(exercice)

	if err != nil {
		panic(err)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: string(exerciceJson),
		},
	})
}
