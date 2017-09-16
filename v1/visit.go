package v1

import (
	"net/http"

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
	beaconID := c.FormValue("beacon_id")
	sendDate := c.FormValue("send_date")
	steps := c.FormValue("steps")
	return c.String(http.StatusOK, "user_id:"+userID+", beacon_id:"+beaconID+", send_date:"+sendDate+", steps:"+steps)
}
