package domain

import "time"

type Lokasi struct {
	ID        int64     `gorm:"primary_key;column:id_lokasi;autoIncrement"`
	Nama      string    `gorm:"column:nama"`
	Alamat    string    `gorm:"column:alamat"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (l *Lokasi) TableName() string {
	return "lokasi"
}
