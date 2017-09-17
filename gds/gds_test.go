package gds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutWeatherData(t *testing.T) {
	weather := Weather{
		Temp: 20.1,
		Wind: 15.5,
		Time: 1505626154,
	}
	err := PutWeatherData(weather)

	assert.Nil(t, err)
}

func TestGetRecentWeather(t *testing.T) {

	currentTime := 1505626154

	w, err := GetRecentWeather(currentTime)

	fmt.Println(w)

	assert.Nil(t, err)
}
