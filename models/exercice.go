package models

import (
	"time"

	"github.com/google/uuid"
)

type Exercice struct {
	Id         uuid.UUID          `json:"id"`
	Type       string             `json:"type"`
	Name       string             `json:"name"`
	Success    bool               `json:"success"`
	Properties ExerciceProperties `json:"properties"`
	CreatedAt  time.Time          `json:"createdAt"`
}
