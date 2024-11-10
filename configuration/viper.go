package configuration

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
    once     sync.Once
    instance *viper.Viper
)

// GetInstance mengembalikan instance singleton dari Viper yang memuat .env
func GetInstance() *viper.Viper {
    once.Do(func() {
        instance = viper.New()

        // Tentukan lokasi dan nama file .env
        instance.SetConfigFile(".env")

        // Membaca file konfigurasi .env
        if err := instance.ReadInConfig(); err != nil {
            fmt.Printf("Error membaca konfigurasi .env: %v", err)
        }

        // Mengatur agar Viper membaca variabel dari environment
        instance.AutomaticEnv()
    })
    return instance
}