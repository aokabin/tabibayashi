package v1

import (
	"net/http"

	"github.com/aokabin/tabibayashi/kvs"
	"github.com/labstack/echo"
)

func Visit(c echo.Context) error {
	// Get name and email
	/*
		user_id: ユーザーのtwitterとかのidになる、例) "twitter#123456789"
		beacon_id: ビーコンのID
		send_date: 送信日時(unix timeとかがいいかもです)
		steps: 前のビーコンからの歩数
	*/

	userID := c.FormValue("user_id")
	vd := kvs.VisitData{}
	vd.BeaconID = c.FormValue("beacon_id")
	vd.SendDate = c.FormValue("send_date")
	vd.Steps = c.FormValue("steps")

	kvs.AddVisitData(userID, &vd)

	return c.String(http.StatusOK, "user_id:"+userID+", beacon_id:"+vd.BeaconID+", send_date:"+vd.SendDate+", steps:"+vd.Steps)
}
