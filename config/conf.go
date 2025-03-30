package config

import "github.com/ericprd/technical-test/utils"

var (
	POSTGRES_HOST,
	POSTGRES_PORT,
	POSTGRES_USER,
	POSTGRES_PASS,
	POSTGRES_DB string
)

func InitConfig() {
	POSTGRES_HOST = utils.GetEnvOrDefault("POSTGRES_HOST", "localhost")
	POSTGRES_PORT = utils.GetEnvOrDefault("POSTGRES_PORT", "5432")
	POSTGRES_USER = utils.GetEnvOrDefault("POSTGRES_USER", "")
	POSTGRES_PASS = utils.GetEnvOrDefault("POSTGRES_PASS", "")
	POSTGRES_DB = utils.GetEnvOrDefault("POSTGRES_DB", "")
}
