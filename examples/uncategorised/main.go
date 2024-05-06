package main

import (
	"context"
	"log"

	"github.com/lowford/policeuk"
)

const (
	serverURL = "https://data.police.uk/api"
)

func main() {
	client, err := policeuk.NewClient(serverURL)

	if err != nil {
		log.Fatalln(err)
	}

	availability, err := client.ListDatasetAvailability(context.TODO())

	if err != nil {
		log.Fatalln(err)
	}

	if len(availability) == 0 {
		log.Fatalln("no dataset availability")
	}

	for _, item := range availability {
		log.Printf("%s - %d police forces have stop-and-search data", item.GetDate().Value, len(item.GetStopAndSearch()))
	}
}
