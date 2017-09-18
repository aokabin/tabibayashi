package music

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aokabin/tabibayashi/gds"
	"github.com/aokabin/tabibayashi/kvs"
	"github.com/aokabin/tabibayashi/storage"
)

func GetMusicFilePath() string {
	// ファイルを一時ファイルに保存しているなら、このあたりに書いていただけると

	return "/path/to/file"
}

func GetMusicData() []byte {
	// byte列で音楽データを保存しているなら、このあたりに

	// 返り値の例
	return bytes.NewBuffer(nil).Bytes()
}

func CreateMusicData(vds []kvs.VisitData, userID string) ([]byte, error) {

	f, err := os.Open("music_test.go")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	url, err := storage.UploadBinaryData(buf, "Goroutine.md")
	if err != nil {
		return nil, err
	}

	music := gds.Music{
		ID:       userID, // 引数で受け取る
		SoundURL: url,
	}

	err = gds.CreateMusicURL(music)
	if err != nil {
		return nil, err
	}

	return buf, nil

}

func CreateMusic(vds []kvs.VisitData, userID string) {

	_, err := CreateMusicData(vds, userID)
	fmt.Println(err)
}
