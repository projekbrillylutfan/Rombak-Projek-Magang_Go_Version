package test

import (
	"testing"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/stretchr/testify/assert"
)



func TestENV(t *testing.T) {
	config := configuration.New("..env")

	assert.Equal(t, "localhost", config.Get("DB_HOST"))
	assert.Equal(t, "admin", config.Get("DB_USER"))
	assert.Equal(t, "localhost", config.Get("DB_HOST"))
	assert.Equal(t, "secret", config.Get("DB_PASSWORD"))
	assert.Equal(t, 8080, config.Get("SERVER_PORT"))
}