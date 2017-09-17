package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	getWeatherSample()
}

func TestGetWeatherData(t *testing.T) {
	getWeatherData()
}

func TestSaveWeatherData(t *testing.T) {
	err := SaveWeatherData()
	assert.Nil(t, err)
}
