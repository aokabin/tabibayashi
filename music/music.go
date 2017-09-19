package music

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aokabin/tabibayashi/gds"
	"github.com/aokabin/tabibayashi/kvs"
	"github.com/aokabin/tabibayashi/storage"
	uuid "github.com/satori/go.uuid"
)

type MusicSource struct {
	Season   string
	Spot     string
	Weather  string
	TimeZone string
	Temp     string
	Wind     string
	Steps    string
}

func GetMusicFilePath() string {
	// ファイルを一時ファイルに保存しているなら、このあたりに書いていただけると

	return "/path/to/file"
}

func GetMusicData() []byte {
	// byte列で音楽データを保存しているなら、このあたりに

	// 返り値の例
	return bytes.NewBuffer(nil).Bytes()
}

// CreateMusicData is 曲を作ってバイナリデータを返す
func CreateMusicData(vds []kvs.VisitData, userID string) ([]byte, error) {

	// f, err := os.Open("../7660.mp3")
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()

	// buf, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	return nil, err
	// }

	// url, err := storage.UploadBinaryData(buf, "Goroutine.md")
	// if err != nil {
	// 	return nil, err
	// }

	mss := []MusicSource{}

	fmt.Println("This is vds")
	fmt.Println(vds)

	for _, vd := range vds {
		ms, err := createMusicSource(vd)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		mss = append(mss, *ms)
	}

	// ここでmssを渡す、[]byteを受け取る
	buf, err := sampleBinaryData(mss)

	fileName := uuid.NewV4().String() + ".wav"
	url, err := storage.UploadBinaryData(buf, fileName)

	music := gds.Music{
		ID:        userID, // 引数で受け取る
		SoundURL:  url,
		CreatedAt: int(time.Now().Unix()),
	}

	err = gds.CreateMusicURL(music)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buf, nil

}

func CreateMusic(vds []kvs.VisitData, userID string) {

	_, err := CreateMusicData(vds, userID)
	fmt.Println(err)
}

func createMusicSource(vd kvs.VisitData) (*MusicSource, error) {
	utime, err := strconv.ParseInt(vd.SendDate, 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	season, err := getSeason(utime)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 一旦ランダム
	spot, err := getSpot(vd.BeaconID)
	weather, err := gds.GetRecentWeather(int(utime))
	// 一旦全部晴れ
	fmt.Println("ok")
	weatherType, err := getWeather(weather.Weather)
	// 一旦全部お昼
	timeZone, err := getTimeZone(utime)
	// 一旦全部暖かい
	temp, err := getTemperature(weather.Temp)
	// 一旦全部弱い、1.0
	wind, err := getWind(weather.Wind)
	steps := vd.Steps

	ms := MusicSource{
		Season:   season,
		Spot:     spot,
		Weather:  weatherType,
		TimeZone: timeZone,
		Temp:     temp,
		Wind:     wind,
		Steps:    steps,
	}

	return &ms, nil
}

func getSeason(utime int64) (string, error) {
	t := time.Unix(utime, 0)
	month := int(t.Month())
	season := ""
	switch month {
	case 3, 4, 5:
		season = "0"
	case 6, 7, 8:
		season = "1"
	case 9, 10, 11:
		season = "2"
	case 12, 1, 2:
		season = "3"
	}
	return season, nil
}

func getSpot(beaconID string) (string, error) {
	// spotList := []string{"100", "200", "300", "301"}
	spotList := []string{"100", "300"}
	return shuffle(spotList), nil

}

func getWeather(weather string) (string, error) {
	return "0", nil
}

func getTimeZone(utime int64) (string, error) {
	return "1", nil
}

func getTemperature(temp float32) (string, error) {
	return "1", nil
}

func getWind(wind float32) (string, error) {
	return "1.0", nil
}

func sampleBinaryData(mss []MusicSource) ([]byte, error) {
	fmt.Println("Creating Music...")
	f, err := os.Open("/Users/kd/go/src/github.com/aokabin/tabibayashi/sample.wav")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Music was created!...")
	return buf, nil

}

func shuffle(list []string) string {
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i) // 0～(i-1) の乱数発生
		list[i-1], list[j] = list[j], list[i-1]
	}

	return list[0]
}
