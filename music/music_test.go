package music

import (
	"fmt"
	"testing"
	"time"

	"github.com/aokabin/tabibayashi/kvs"
)

func TestGetMusicFilePath(t *testing.T) {
	str := GetMusicFilePath()
	fmt.Println(str)
}

func TestGetMusicData(t *testing.T) {
	buf := GetMusicData()
	fmt.Println(buf)
}

func TestCreateMuisicData(t *testing.T) {

	v1 := kvs.VisitData{
		BeaconID: "hogehoge",
		SendDate: "0000000",
		Steps:    "1000",
	}
	v2 := kvs.VisitData{
		BeaconID: "hoge",
		SendDate: "0000",
		Steps:    "10",
	}

	vds := []kvs.VisitData{v1, v2}
	fmt.Println(vds)

	go CreateMusic(vds, "sample_id")
	time.Sleep(30000000000)

	// assert.Nil(t, err)

}
