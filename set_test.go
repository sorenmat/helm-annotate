package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	x := []string{"test=yes", "test1=no"}
	anno, err := toMap(x)
	assert.NoError(t, err)
	assert.Equal(t, "yes", anno["test"])
}

func TestToMapMultipleEquals(t *testing.T) {
	x := []string{"test=yes", "test1=Joe=Man"}
	anno, err := toMap(x)
	assert.NoError(t, err)
	assert.Equal(t, "Joe=Man", anno["test1"])
}

func TestToMapLen(t *testing.T) {
	x := []string{"test=yes", "test1=Joe=Man", "cat=hat"}
	anno, err := toMap(x)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(anno))
}

func TestToMapMissingEqual(t *testing.T) {
	x := []string{"test", "test1=Joe=Man", "cat=hat"}
	_, err := toMap(x)
	assert.Error(t, err)
}
