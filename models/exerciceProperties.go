package models

import (
	"log"

	"github.com/google/uuid"
	"github.com/matthieurobert/workout-bot/config"
)

type ExerciceProperties struct {
	Id       uuid.UUID `json:"id"`
	Distance float32   `json:"distance"`
	Length   int       `json:"lenth"`
	Rep      int       `json:"rep"`
	Series   []int     `json:"series"`
}

func CreateExerciceProperties(ep *ExerciceProperties) {
	_, err := config.POSTGRES.Db.Model(ep).Insert()

	if err != nil {
		log.Panicln(err)
	}
}
