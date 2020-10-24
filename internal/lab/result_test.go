package lab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/comptag/bobcat-lamp/internal/lab"
)

func TestResultGetter(t *testing.T) {
	result := lab.MakeResult("abc", "bob", "123", false)

	assert.Equal(t, result.Id(), "abc")
}
