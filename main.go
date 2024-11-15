package main

import (
	"fmt"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
)

func main() {
    config := configuration.New()

    // Mengakses konfigurasi menggunakan Viper singleton
    dbHost := config.Get("DB_HOST")
    dbUser := config.Get("DB_USER")
    dbPassword := config.Get("DB_PASSWORD")
    serverPort := config.Get("SERVER_PORT")

    fmt.Println("DB_HOST:", dbHost)
    fmt.Println("DB_USER:", dbUser)
    fmt.Println("DB_PASSWORD:", dbPassword)
    fmt.Println("SERVER_PORT:", serverPort)

    // Contoh penggunaan jika variabel tidak ditemukan
}