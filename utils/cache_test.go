// Retrieve existing cache value successfully
package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCacheExistingKey(t *testing.T) {
	key := "existingKey"
	expectedValue := "value"

	// Set the cache value
	SetCache(key, expectedValue)

	// Retrieve the cache value
	value, err := GetCache(key)

	// Assert no error and correct value
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
}
