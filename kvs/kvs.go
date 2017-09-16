package kvs

import (
	"bytes"
	"encoding/gob"

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
	Connection() //TODO: 絶対もっといい方法ある
	return c.RPush(key, value).Result()
}

func getLastListData(key string) (*VisitData, error) {
	Connection() //TODO: 絶対もっといい方法ある
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
