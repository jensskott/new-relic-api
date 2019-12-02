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

	config := make(map[string]string)

	config["recipients"] = "test@massive.co"
	config["include_json_attachment"] = "true"

	body := new_relic_api.ChannelsPayload{Channel: new_relic_api.Channel{
		Name:          "Test",
		Type:          "email",
		Configuration: config,
	}}

	resp, err := c.CreateAlertsChannels(body)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(resp)
}
