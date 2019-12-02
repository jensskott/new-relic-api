package convert_test

import (
	"github.com/stretchr/testify/assert"
	"stash.massiveinteractive.com/to/new-relic-api/lib/convert"
	"testing"
)

func TestStr(t *testing.T) {
	testStr := convert.Str("test")

	assert.Equal(t, "test", *testStr)
}
