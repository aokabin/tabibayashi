package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisHost(t *testing.T) {
	sampleENV := "localhost"
	os.Setenv("TBS_REDIS_HOST", sampleENV)
	reload()
	assert.Equal(t, RedisHost(), sampleENV)
}
