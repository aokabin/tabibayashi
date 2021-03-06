package gds

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/aokabin/tabibayashi/config"
	"google.golang.org/api/iterator"
)

var (
	c    *datastore.Client
	ctx  context.Context
	kind string
)

type Weather struct {
	Weather string  `datastore:",noindex"`
	Temp    float32 `datastore:",noindex"`
	Wind    float32 `datastore:",noindex"`
	Time    int
}

type Beacon struct {
	ID         string
	MajorValue string
	MinorValue string
	CreatedAt  int
}

type Music struct {
	ID        string
	SoundURL  string
	CreatedAt int
}

func init() {
	ctx = context.Background()
	kind = "Weather"
	Connection()
}

func Connection() {

	// Set your Google Cloud Platform project ID.
	projectID := config.ProjectID()

	// Creates a client.
	var err error
	c, err = datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// インターフェース的にはデータは型を決めたくない気持ちがある

func PutWeatherData(w Weather) error {
	ctx := context.Background()
	weatherKey := datastore.IncompleteKey(kind, nil)

	_, err := c.Put(ctx, weatherKey, &w)

	return err
}

func GetRecentWeather(time int) (*Weather, error) {
	query := datastore.NewQuery(kind).Order("-Time").Filter("Time <= ", time).Limit(1)
	it := c.Run(ctx, query)
	var beforeWeather Weather
	for {
		_, err := it.Next(&beforeWeather)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	fmt.Printf("Before Weather %q\n", beforeWeather.Time)

	query = datastore.NewQuery(kind).Order("Time").Filter("Time >= ", time).Limit(1)
	it = c.Run(ctx, query)
	var afterWeather Weather
	for {
		_, err := it.Next(&afterWeather)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	fmt.Printf("After Weather %q\n", afterWeather.Time)

	subBef := time - beforeWeather.Time
	subAft := afterWeather.Time - time

	if subBef > subAft {
		return &afterWeather, nil
	}

	return &beforeWeather, nil

}

func GetAllBeacons() ([]Beacon, error) {
	query := datastore.NewQuery("Beacon")
	it := c.Run(ctx, query)
	beacons := []Beacon{}
	for {
		var beacon Beacon
		_, err := it.Next(&beacon)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		beacons = append(beacons, beacon)

	}

	return beacons, nil
}

func CreateBeacon(b Beacon) error {
	ctx := context.Background()
	beaconKey := datastore.IncompleteKey("Beacon", nil)

	_, err := c.Put(ctx, beaconKey, &b)

	return err
}

func CreateMusicURL(m Music) error {
	musicKey := datastore.IncompleteKey("Music", nil)

	_, err := c.Put(ctx, musicKey, &m)

	return err
}

func GetMusicURL(userID string) (*Music, error) {
	query := datastore.NewQuery("Music").Order("-CreatedAt").Limit(1)
	it := c.Run(ctx, query)
	var music Music
	for {
		_, err := it.Next(&music)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return &music, nil
}
