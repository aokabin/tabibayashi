package kvs

import (
	"bytes"
	"encoding/gob"

	"github.com/motemen/ghq/utils"

	"github.com/aokabin/tabibayashi/config"
	"github.com/go-redis/redis"
)

var (
	c *redis.Client
)

type VisitData struct {
	BeaconID string
	SendDate string
	Steps    string
}

func init() {
	err := Connection()
	utils.PanicIf(err)
}

func Connection() error {
	host := config.RedisHost()
	port := config.RedisPort()

	c = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	return err
}

func addList(key string, value []byte) (int64, error) {
	return c.RPush(key, value).Result()
}

func getLastListData(key string) (*VisitData, error) {
	binVD, err := c.LIndex(key, -1).Result()
	buf := bytes.NewBuffer([]byte(binVD))
	var vd VisitData
	gob.NewDecoder(buf).Decode(&vd)
	return &vd, err
}

func AddVisitData(key string, vd *VisitData) (int64, error) {
	buf := bytes.NewBuffer(nil)
	gob.NewEncoder(buf).Encode(&vd)
	return addList(key, buf.Bytes())
}

func GetLastVisitData(key string) (*VisitData, error) {
	vd, err := getLastListData(key)
	return vd, err
}

func GetAllVisitData(key string) ([]VisitData, error) {
	var vds []VisitData
	binVDs, err := c.LRange(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, d := range binVDs {
		buf := bytes.NewBuffer([]byte(d))
		var vd VisitData
		gob.NewDecoder(buf).Decode(&vd)
		vds = append(vds, vd)
	}

	return vds, nil
}

func RemoveVisitData(key string) error {
	err := c.Del(key).Err()
	return err
}
