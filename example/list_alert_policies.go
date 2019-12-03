package main

import (
	"fmt"
	"log"
	"os"

	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"
)

func main() {
	// use go run get_alert_channels.go yourapikeyhere
	c := new_relic_api.New(os.Args[1])

	resp, err := c.ListAlertsPolicies("test-api-client", "false")
	if err != nil {
		log.Fatal(err.Error())
	}

	//resp, err := c.ListAlertsPolicies("", "")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	fmt.Println(resp)
}
