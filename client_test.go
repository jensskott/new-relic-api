package new_relic_api_test

import (
	"github.com/stretchr/testify/assert"
	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := new_relic_api.New("1234")

	assert.NoError(t, err)
	assert.Equal(t, "1234", *c.ApiKey)
}
