package app

import (
	"log"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdminUser(DB *gorm.DB) {
	// Cek apakah user dengan username 'test' sudah ada
	var user *domain.User
	result := DB.Where("username = ?", "test").First(&user)
	if result.RowsAffected == 0 {
		// Hash password sebelum disimpan
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		// Buat user baru
		newUser := &domain.User{
			Nama:     "test",
			Jabatan:  "test",
			Username: "test",
			Email:    "test@example.com",
			Password: string(hashedPassword),
			Role:     "ADMIN",
		}

		if err := DB.Create(&newUser).Error; err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
		log.Println("Admin user created successfully")
	} else {
		log.Println("Admin user already exists")
	}
}