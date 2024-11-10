package main

import (
	"fmt"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
)

func main() {
	viper := configuration.GetInstance()

    // Mengakses konfigurasi menggunakan Viper singleton
    dbHost := viper.GetString("DB_HOST")
    dbUser := viper.GetString("DB_USER")
    dbPassword := viper.GetString("DB_PASSWORD")
    serverPort := viper.GetInt("SERVER_PORT")

    fmt.Println("DB_HOST:", dbHost)
    fmt.Println("DB_USER:", dbUser)
    fmt.Println("DB_PASSWORD:", dbPassword)
    fmt.Println("SERVER_PORT:", serverPort)

    // Contoh penggunaan jika variabel tidak ditemukan
    if serverPort == 0 {
        fmt.Println("SERVER_PORT belum diatur di file .env")
    }
}