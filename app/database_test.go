package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBConnection(t *testing.T) {
	db := GetDBConnection()
	assert.NotEmpty(t, db)
}
