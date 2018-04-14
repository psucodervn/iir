package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestCase(t *testing.T) {
	should := assert.New(t)

	testCase := NewTestCase("input", "output")
	should.NotNil(testCase)
	should.Equal(testCase.Input, "input")
	should.Equal(testCase.Output, "output")
}

func TestDefaultTaskImplementsTask(t *testing.T) {
	should := assert.New(t)
	should.Implements((*Task)(nil), new(DefaultTask))
}