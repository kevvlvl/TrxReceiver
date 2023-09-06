package env

import (
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func ApiPort() string {
	return envStr("API_PORT")
}

func RedisHost() string {
	return envStr("REDIS_HOST")
}

func RedisPass() string {
	return envStr("REDIS_PASS")
}

func RedisPort() int64 {
	return envInt("REDIS_PORT", 10, 0)
}

func envStr(k string) string {
	return os.Getenv(k)
}

func envInt(k string, base int, bitSize int) int64 {

	v, err := strconv.ParseInt(os.Getenv(k), base, bitSize)

	if err != nil {
		log.Error().Msgf("Error trying to convert %v env var to int", k)
		return -1
	}

	return v
}
