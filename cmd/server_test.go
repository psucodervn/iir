package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicServerImplementsServer(t *testing.T) {
	should := assert.New(t)
	should.Implements((*Server)(nil), new(BasicServer))
}
