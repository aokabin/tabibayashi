package gds

import (
	"fmt"
	"testing"
	"time"

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

func TestGetAllBeacons(t *testing.T) {
	beacons, err := GetAllBeacons()
	fmt.Println(beacons)
	assert.Nil(t, err)
}

func TestCreateBeacon(t *testing.T) {
	beacon := Beacon{
		ID:         "zzzzzzz",
		MajorValue: "1",
		MinorValue: "0",
		CreatedAt:  int(time.Now().Unix()),
	}

	err := CreateBeacon(beacon)

	assert.Nil(t, err)

}
