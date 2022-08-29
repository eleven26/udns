//go:build integration

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIp(t *testing.T) {
	_, err := Ip()

	assert.Nil(t, err)
}
