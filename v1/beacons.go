package v1

import (
	"fmt"
	"net/http"

	"github.com/aokabin/tabibayashi/gds"
	"github.com/labstack/echo"
)

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
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, beacons)
}
