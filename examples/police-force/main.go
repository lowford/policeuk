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

	forces, err := client.ListPoliceForces(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	for _, force := range forces {
		res, err := client.GetPoliceForce(context.Background(), policeuk.GetPoliceForceParams{
			ForceId: force.GetID().Value,
		})

		if err != nil {
			log.Fatalln(err)
		}

		switch info := res.(type) {
		case *policeuk.PoliceForce:
			log.Printf(
				"%s (%s) - Tel %s - Url %v - Engagement Methods %d",
				info.GetName().Value,
				info.GetID().Value,
				info.GetTelephone().Value,
				info.GetURL().Value,
				len(info.GetEngagementMethods().Or(make([]policeuk.PoliceForceEngagementMethodsItem, 0))),
			)
		case *policeuk.GetPoliceForceNotFound:
			log.Fatalln("police force not found")
		}
	}
}
