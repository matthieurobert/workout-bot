package config

import (
	"os"
	"strconv"
)

type Env struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func (e *Env) initEnv() {
	var err error

	e.PostgresHost = os.Getenv("POSTGRES_HOST")
	e.PostgresPort, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	e.PostgresUser = os.Getenv("POSTGRES_USER")
	e.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	e.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")

	if err != nil {
		panic(err)
	}
}
