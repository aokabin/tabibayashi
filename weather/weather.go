package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/aokabin/tabibayashi/gds"
)

type CurrentObservation struct {
	Temp    float32 `json:"temp_c"`
	WindKph float32 `json:"wind_kph"`
	Time    string  `json:"observation_epoch"`
}

type WeatherJSON struct {
	CurrentObservation CurrentObservation `json:"current_observation"`
}

/** JSONデコード用に構造体定義 */

func SaveWeatherData() error {
	w, err := getWeatherData()
	if err != nil {
		fmt.Println(err)
		return err
	}
	time, err := strconv.Atoi(w.CurrentObservation.Time)
	if err != nil {
		fmt.Println(err)
		return err
	}
	weather := gds.Weather{
		Temp: w.CurrentObservation.Temp,
		Wind: w.CurrentObservation.WindKph,
		Time: time,
	}
	err = gds.PutWeatherData(weather)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getWeatherData() (*WeatherJSON, error) {
	resp, err := http.Get("http://api.wunderground.com/api/cc4e59b34bd83b1f/conditions/q/JP/Kyoto.json")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var weather WeatherJSON
	err = json.Unmarshal([]byte(string(body)), &weather)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &weather, nil
}

func getWeatherSample() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("kyoto.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var weather WeatherJSON
	if err := json.Unmarshal(bytes, &weather); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	fmt.Printf("%v\n", weather.CurrentObservation.Time)
}
