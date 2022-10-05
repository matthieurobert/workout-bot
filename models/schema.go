package models

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func CreateSchema(db *pg.DB) {
	models := []interface{}{
		(*Exercice)(nil),
		(*ExerciceProperties)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
