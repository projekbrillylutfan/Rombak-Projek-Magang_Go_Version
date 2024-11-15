package domain

import "time"

type User struct {
	ID        int64 `gorm:"primaryKey"`
	Nama      string ``
	Jabatan   string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}

func (u *User) TableName() string {
	return "users"
}
