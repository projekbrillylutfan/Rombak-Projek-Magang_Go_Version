package test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
)



func TestMain(m *testing.M) {
	// Pastikan path ke .env relatif terhadap root proyek
	err := godotenv.Load("../.env")
	exception.PanicLogging(err)
	os.Exit(m.Run())
}

func TestConfig(t *testing.T) {
	config := configuration.New()
	dbHost := config.Get("DB_HOST")
	if dbHost == "" {
		t.Fatal("Expected DB_HOST to be set in environment")
	}
}