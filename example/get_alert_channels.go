package main

import (
	"fmt"
	"os"

	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"
)

func main() {
	// use go run get_alert_channels.go yourapikeyhere
	c, _ := new_relic_api.New(os.Args[1])
	resp, err := c.ListAlertChannels()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)
}
