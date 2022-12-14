package commands

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/matthieurobert/workout-bot/models"
)

var integerMinRepValue = 0

var Exercice = &discordgo.ApplicationCommand{
	Name:        "exercice",
	Description: "Add an exercice",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Description: "Name of the exercice",
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Leg press1",
					Value: "Leg press",
				},
				{
					Name:  "Adductor1",
					Value: "Adductor",
				},
			},
		},
		{
			// Type:        discordgo.ApplicationCommandOptionInteger,
			Type:        discordgo.ApplicationCommandOptionNumber,
			Name:        "rep",
			Description: "Number of rep",
			Required:    true,
			// MinValue:    &integerOptionMinValue,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "series",
			Description: "Series of the exercice",
			Required:    true,
		},
	},
}

func ExerciceHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var name string
	var rep int
	var series []int
	id := uuid.New()

	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if opt, ok := optionMap["name"]; ok {
		// msg += opt.StringValue()
		name = opt.StringValue()
	}

	if opt, ok := optionMap["rep"]; ok {
		// msg += " " + fmt.Sprintf("%g", opt.FloatValue())
		// msg += " " + string(rune(opt.IntValue()))
		rep = int(opt.FloatValue())
	}

	if opt, ok := optionMap["series"]; ok {
		// for _, rune := range opt.StringValue() {
		// 	fmt.Println("Rune: " + string(rune))
		// 	series = append(series, uint32(rune))
		// }

		stringArray := strings.Fields(opt.StringValue())
		for _, rune := range stringArray {
			str, _ := strconv.Atoi(rune)
			series = append(series, str)
		}
	}

	properties := models.ExerciceProperties{
		Id:     id,
		Rep:    rep,
		Series: series,
	}

	models.CreateExerciceProperties(&properties)

	exercice := models.Exercice{
		Id:         id,
		Name:       name,
		Type:       "rep",
		Success:    true,
		Properties: &properties,
		CreatedAt:  time.Now(),
	}

	models.CreateExercice(&exercice)

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
