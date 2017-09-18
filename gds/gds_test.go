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

func TestCreateMusicURL(t *testing.T) {
	music := Music{
		ID:       "xxx",
		SoundURL: "https://storage.googleapis.com/tabibayashi-musics/a4a902d7-24c5-4b4a-a800-226dcc260600Goroutine.md",
		// SoundURL:  "https://drive.google.com/open?id=0B2XteJUUmo3Jd3NwN05UNlNObEE",
		CreatedAt: int(time.Now().Unix()),
	}

	err := CreateMusicURL(music)

	assert.Nil(t, err)

}

func TestGetMusicURL(t *testing.T) {
	music, err := GetMusicURL("xxx")
	assert.Nil(t, err)
	fmt.Println(music)
}
