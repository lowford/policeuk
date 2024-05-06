package main

import (
	"context"
	"log"

	"github.com/lowford/policeuk"
)

func main() {
	client, err := policeuk.NewClient(policeuk.ServerURL)

	if err != nil {
		log.Fatalln(err)
	}

	availability, err := client.GetDatasetAvailability(context.TODO())

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
