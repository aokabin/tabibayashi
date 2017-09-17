package kvs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddVisitData(t *testing.T) {
	key := "user_sample"
	vd := VisitData{
		BeaconID: "10",
		SendDate: "100000",
		Steps:    "4000",
	}

	val, err := AddVisitData(key, &vd)

	fmt.Println(val)

	assert.Nil(t, err)
	// assert.
}

func TestGetLastVisitData(t *testing.T) {
	key := "x"
	vd := VisitData{
		BeaconID: "10",
		SendDate: "100000",
		Steps:    "5000",
	}

	val, err := AddVisitData(key, &vd)
	fmt.Println(val)
	assert.Nil(t, err)

	lastVD, err := GetLastVisitData(key)

	fmt.Println(vd)
	fmt.Println(lastVD)

	assert.Nil(t, err)
	assert.NotNil(t, vd)
	assert.Equal(t, vd, *lastVD)
}

func TestGetAllVisitData(t *testing.T) {
	vds, err := GetAllVisitData("x")
	assert.Nil(t, err)

	fmt.Println(vds)
}
