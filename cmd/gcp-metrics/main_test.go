package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	l := map[string]string{
		"action": "test",
		"actor":  "test",
	}

	_, _, err := count("", "test", "1", l)
	assert.Error(t, err)

	_, _, err = count("test", "", "1", l)
	assert.Error(t, err)

	_, _, err = count("test", "test", "", l)
	assert.Error(t, err)

	m, v, err := count("test", "test-metric", "1", l)
	assert.NoError(t, err)
	assert.Equal(t, "custom.googleapis.com/test-metric", m)
	assert.Equal(t, int64(1), v)
}
