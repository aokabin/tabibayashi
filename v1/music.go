package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aokabin/tabibayashi/gds"
	"github.com/aokabin/tabibayashi/kvs"
	"github.com/aokabin/tabibayashi/music"
	"github.com/labstack/echo"
)

type RecievedData struct {
	VisitData []VisitData `json:"visits"`
}

type VisitData struct {
	BeaconID string `json:"beacon_id"`
	SendDate string `json:"send_date"`
	Steps    string `json:"steps"`
}

type Data struct {
	URL string
}

func CreateMusic(c echo.Context) error {
	userID := c.FormValue("user_id")
	var vd []VisitData
	visits := c.FormValue("visits")
	err := json.Unmarshal([]byte(visits), &vd)

	fmt.Println(vd)

	for _, v := range vd {
		visitData := kvs.VisitData{
			BeaconID: v.BeaconID,
			SendDate: v.SendDate,
			Steps:    v.Steps,
		}
		_, err := kvs.AddVisitData(userID, &visitData)
		if err != nil {
			return err
		}
	}

	vds, err := kvs.GetAllVisitData(userID)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: このあたりで作曲の関数を呼ぶ、goroutineかな
	go music.CreateMusic(vds, userID)
	for _, vd := range vds {
		fmt.Println(vd)
	}

	// 作曲が正常に終わったらキーを削除
	kvs.RemoveVisitData(userID)

	return c.String(http.StatusAccepted, "Accepted")
}

func GetMusicURL(c echo.Context) error {
	userID := c.FormValue("user_id")
	fmt.Println(userID)
	/*
		1. ユーザーの情報をdatastoreに取りに行く
		2. もしurlの欄がなければ、まだできていないので、まだできてないレスポンスを返す
		3. あれば、そのurlを返す
	*/

	music, err := gds.GetMusicURL(userID)
	if err != nil {
		return err
	}

	data := Data{
		URL: music.SoundURL,
	}

	return c.JSON(http.StatusOK, data)

}
