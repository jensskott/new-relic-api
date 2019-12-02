package new_relic_api_test

import (
	"testing"

	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := new_relic_api.New("1234")

	assert.Equal(t, "1234", *c.ApiKey)
}
