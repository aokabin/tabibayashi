package music

import (
	"fmt"
	"testing"
)

func TestGetMusicFilePath(t *testing.T) {
	str := GetMusicFilePath()
	fmt.Println(str)
}

func TestGetMusicData(t *testing.T) {
	buf := GetMusicData()
	fmt.Println(buf)
}
