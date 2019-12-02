package main

import (
	"fmt"
	"log"
	"os"

	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"
)

func main() {
	c := new_relic_api.New(os.Args[1])
	resp, err := c.ListAlertsConditions("73380")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(resp)
}
