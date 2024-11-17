package domain

import "time"

type User struct {
	ID              int64     `gorm:"primary_key;column:id_user;autoIncrement"`
	Nama            string    `gorm:"column:nama"`
	Jabatan         string    `gorm:"column:jabatan"`
	Username        string    `gorm:"column:username"`
	Email           string    `gorm:"column:email"`
	Password        string    `gorm:"column:password"`
	Role            string    `gorm:"column:role"`
	IsEmailVerified bool      `gorm:"column:is_email_verified"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt        time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeleteAt        time.Time `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
