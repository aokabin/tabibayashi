package music

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	"bufio"
	"bytes"
	"flag"

	"gopkg.in/pkg/profile.v1"

	"github.com/go-mix/mix"
	"github.com/go-mix/mix/bind"
	"github.com/go-mix/mix/bind/spec"
	"github.com/viert/lame"
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
	spotList := []string{"100", "200", "300", "301"}
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
	pattern := []string{}
	rtnbuf := new(bytes.Buffer)
	loops := len(mss)
	pattern = selectPattern(mss)
	rtnbuf = joinMusic(loops, pattern)
	return nil, rtnbuf
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
func selectPattern(mss []MusicSource) []string {
	pattern := []string{}
	phraseName := ""
	randomNum := ""
	melody := ""
	chord := ""
	bass := ""
	pad := ""
	drums := ""
	perc := ""
	for _, d := range mss {
		
		//歩数の1桁目からランダマイズ用の値を取得
		randomNum = string(d.Steps[(len(d.Steps) - 1):len(d.Steps)]) 
		
		num, err := strconv.Atoi(randomNum)
		if err != nil {
			log.Fatal(err)
		}

		if (num % 2) == 0 {
			randomNum = "0"  //偶数
		} else {
			randomNum = "1" //奇数
		}

		phraseName += d.Season + "_"
		phraseName += d.Spot + "_"
		phraseName += d.Weather + "_"

		melody += "melody/" + phraseName + randomNum + ".wav"
		chord += "chord/" + phraseName + randomNum + ".wav"
		bass += "bass/" + phraseName + randomNum + ".wav"

		phraseName += d.TimeZone + "_"
		pad += "pad/" + phraseName + randomNum + ".wav"

		phraseName += d.Temp + "_"

		drums += "drums/" + phraseName + randomNum + ".wav"
		perc += "perc/" + phraseName + randomNum + ".wav"

		pattern = append(pattern, melody)
		pattern = append(pattern, chord)
		pattern = append(pattern, bass)
		pattern = append(pattern, pad)
		pattern = append(pattern, drums)
		pattern = append(pattern, perc)

		melody = ""
		chord = ""
		bass = ""
		pad = ""
		drums = ""
		perc = ""
		phraseName = ""
	}

	return pattern
}


var profileMode string
var loader, out string
prefix := "sound/"
bpm := 120
step := time.Minute / time.Duration(bpm/8)
sampleHz := float64(44100)
specs := spec.AudioSpec{
	Freq:     sampleHz,
	Format:   spec.AudioS16,
	Channels: 2,
}

// command-line arguments
flag.StringVar(&out, "out", "null", "playback binding [null] _OR_ [wav] for direct stdout (e.g. >file or |aplay)")
flag.StringVar(&profileMode, "profile", "", "enable profiling [cpu, mem, block]")
flag.StringVar(&loader, "loader", "wav", "input loading interface [wav, sox]")
flag.Parse()

// CPU/Memory/Block profiling
if len(profileMode) > 0 {
	out = "null" // TODO: evaluate whether profiling is actually working
	switch profileMode {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.MemProfileRate(4096)).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	default:
		// do nothing
	}
}

// configure mix
bind.UseOutputString(out)
bind.UseLoaderString(loader)
defer mix.Teardown()
mix.Configure(specs)
mix.SetSoundsPath(prefix)

// setup the music
t := 1 * time.Second // buffer before music for 1 second

for n := 0; n < loops; n++ {
	for s := 0; s < 6; s++ {
		mix.SetFire(
			pattern[s+n*6], t+time.Duration(s), 0, 1.0, 0)
	}
	t += time.Duration(6) * step
}
t += 4 * time.Second // buffer after music for 4 seconds

//var data io.Writer
buf := new(bytes.Buffer)
if bind.IsDirectOutput() {

	mix.Debug(true)
	writer := bufio.NewWriter(buf)
	mix.OutputStart(t, writer)
	for p := time.Duration(0); p <= t; p += t / 4 {
		mix.OutputContinueTo(p)
	}
	mix.OutputClose()

} else {
	mix.Debug(true)
	mix.StartAt(time.Now().Add(1 * time.Second))
	fmt.Printf("Mix: 808 Example - pid:%v playback:%v spec:%v\n", os.Getpid(), out, specs)
	for mix.FireCount() > 0 {
		time.Sleep(1 * time.Second)
	}
}

rtnbuf := new(bytes.Buffer)
rtnbuf = convertMp3(buf)

return rtnbuf

}
func convertMp3(buf *bytes.Buffer) *bytes.Buffer {
	
		reader := bufio.NewReader(buf)
	
		of, err := os.Create("output10.mp3")
		if err != nil {
			panic(err)
		}
		defer of.Close()
		rtnbuf := new(bytes.Buffer)
		wr := lame.NewWriter(rtnbuf)
		//wr := lame.NewWriter(of)
		wr.Encoder.SetBitrate(320)
		wr.Encoder.SetQuality(1)
	
		// IMPORTANT!
		wr.Encoder.InitParams()
	
		reader.WriteTo(wr)
		return rtnbuf
	}
	

func shuffle(list []string) string {
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i) // 0～(i-1) の乱数発生
		list[i-1], list[j] = list[j], list[i-1]
	}

	return list[0]
}
