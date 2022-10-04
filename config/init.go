package config

var ENV = Env{}

var POSTGRES = &Postgres{}

func InitConfig() {
	ENV.initEnv()

	POSTGRES.connectToDb(&ENV)

}
