package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/matthieurobert/workout-bot/config"
)

type Exercice struct {
	Id         uuid.UUID           `json:"id"`
	Type       string              `json:"type"`
	Name       string              `json:"name"`
	Success    bool                `json:"success"`
	Properties *ExerciceProperties `json:"properties" pg:"rel:has-one"`
	CreatedAt  time.Time           `json:"createdAt"`
}

func CreateExercice(e *Exercice) {
	_, err := config.POSTGRES.Db.Model(e).Insert()

	if err != nil {
		log.Panicln(err)
	}
}
