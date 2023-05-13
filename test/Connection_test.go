package test

import (
	"testing"

	"github.com/juandaantoniusapakpahan/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestCleanUpd(t *testing.T) {
	request, result := simple.InitializedConnectoin("Go.go")

	assert.NotNil(t, request)
	result()
}
