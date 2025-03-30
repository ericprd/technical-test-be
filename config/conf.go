package config

import (
	"fmt"

	"github.com/ericprd/technical-test/utils"
)

var (
	POSTGRES_HOST,
	POSTGRES_PORT,
	POSTGRES_USER,
	POSTGRES_PASS,
	POSTGRES_DB,
	REDIS_HOST,
	REDIS_PASS string

	REDIS_DB int

	REDIS_DISABLE_ID bool
)

func InitConfig() {
	POSTGRES_HOST = utils.GetEnvOrDefault("POSTGRES_HOST", "localhost")
	POSTGRES_PORT = utils.GetEnvOrDefault("POSTGRES_PORT", "5432")
	POSTGRES_USER = utils.GetEnvOrDefault("POSTGRES_USER", "")
	POSTGRES_PASS = utils.GetEnvOrDefault("POSTGRES_PASS", "")
	POSTGRES_DB = utils.GetEnvOrDefault("POSTGRES_DB", "")

	REDIS_HOST = fmt.Sprintf(
		"%s:%d",
		utils.GetEnvOrDefault("REDIS_HOST", "localhost"),
		utils.GetEnvOrDefault("REDIS_PORT", 6379),
	)

	REDIS_PASS = utils.GetEnvOrDefault("REDIS_PASS", "")
	REDIS_DB = utils.GetEnvOrDefault("REDIS_DB", 0)

	REDIS_DISABLE_ID = utils.GetEnvOrDefault("REDIS_DISABLE_ID", false)
}
