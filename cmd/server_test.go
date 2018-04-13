package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementServer(t *testing.T) {
	as := assert.New(t)

	as.Implements((*Server)(nil), new(BasicServer))
}
