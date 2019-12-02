package main

import (
	"fmt"
	"os"

	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"
)

func main() {
	// use go run get_alert_channels.go yourapikeyhere
	c := new_relic_api.New(os.Args[1])

	config := make(map[string]string)

	config["recipients"] = "test@test.test"
	config["include_json_attachment"] = "true"

	body := new_relic_api.Payload{Channel: new_relic_api.Channel{
		Name:          "Test",
		Type:          "email",
		Configuration: config,
	}}

	resp, err := c.CreateAlertChannels(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)
}
