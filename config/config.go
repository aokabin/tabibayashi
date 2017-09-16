package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	RedisHost string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort string `envconfig:"REDIS_PORT" default:"6379"`
}

var (
	c Configuration
)

const (
	prefix = "TBS"
)

func init() {
	envconfig.MustProcess(prefix, &c)
}

func reload() {
	envconfig.Process(prefix, &c)
}

func RedisHost() string {
	return c.RedisHost
}

func RedisPort() string {
	return c.RedisPort
}
