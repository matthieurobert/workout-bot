package config

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-pg/pg/v10"
)

type Postgres struct {
	Db *pg.DB
}

func (ps *Postgres) connectToDb(env *Env) {
	ps.Db = pg.Connect(&pg.Options{
		Addr: env.PostgresHost + ":" + strconv.Itoa(env.PostgresPort),
		// Addr:     "postgres:5432",
		User:     env.PostgresUser,
		Password: env.PostgresPassword,
		Database: env.PostgresDatabase,
	})

	fmt.Println(ps.Db.String())
	fmt.Println(env.PostgresDatabase)

	ctx := context.Background()

	if err := ps.Db.Ping(ctx); err != nil {
		panic(err)
	}
}
