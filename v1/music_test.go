package v1

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	j := `[{"beacon_id":"yyyyyyy","send_date":"1505704785","steps":"100"}]`
	var vd []VisitData
	err := json.Unmarshal([]byte(j), &vd)

	fmt.Println(vd)
	assert.Nil(t, err)
}
