package domain

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;column:id_user;autoIncrement"`
	Nama      string    `gorm:"column:nama"`
	Jabatan   string    `gorm:"column:jabatan"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeleteAt  time.Time `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
