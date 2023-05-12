package test

import (
	"testing"

	"github.com/juandaantoniusapakpahan/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleDependencyError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleDependencySuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
