package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	ProjectID  string `envconfig:"PROJECT_ID" default:"kyotohack19-team-a"`
	RedisHost  string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort  string `envconfig:"REDIS_PORT" default:"6379"`
	ServerPort string `default:"1323"`
	BucketName string `default:"tabibayashi-musics"`
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

func ProjectID() string {
	return c.ProjectID
}

func RedisHost() string {
	return c.RedisHost
}

func RedisPort() string {
	return c.RedisPort
}

func ServerPort() string {
	return c.ServerPort
}

func BucketName() string {
	return c.BucketName
}
