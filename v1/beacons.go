package v1

import (
	"net/http"

	"github.com/aokabin/tabibayashi/gds"
	"github.com/labstack/echo"
)

type Beacons struct {
	Info []Beacon `json:"info"`
}

type Beacon struct {
	ID         string `json:"id"`
	MajorValue string `json:"major_value"`
	MinorValue string `json:"minor_value"`
}

func GetBeacons(c echo.Context) error {

	getBeacons, err := gds.GetAllBeacons()
	beacons := []Beacon{}
	for _, b := range getBeacons {
		beacon := Beacon{
			ID:         b.ID,
			MajorValue: b.MajorValue,
			MinorValue: b.MinorValue,
		}
		beacons = append(beacons, beacon)
	}
	if err != nil {
		return err
	}

	info := Beacons{
		Info: beacons,
	}

	return c.JSON(http.StatusOK, info)
}
