package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	_, err := GenerateToken(1)

	assert.NoError(t, err)
}
