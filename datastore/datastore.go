package datastore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/aokabin/tabibayashi/config"
)

var (
	c *datastore.Client
)

type Weather struct {
	Temp int `datastore:",noindex"`
	Wind int `datastore:",noindex"`
	Time int
}

func init() {
	Connection()
}

func Connection() {
	ctx := context.Background()

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

func PutWeatherData() {
	ctx := context.Background()
	// Sets the kind for the new entity.
	kind := "Weather"
	// Creates a Key instance.
	weatherKey := datastore.IncompleteKey(kind, nil)

	fmt.Println(weatherKey)

	// Creates a Task instance.
	weather := Weather{
		Temp: 20,
		Wind: 15,
		Time: 1505626154,
	}

	// Saves the new entity.
	if _, err := c.Put(ctx, weatherKey, &weather); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", weatherKey, weather.Time)
}
