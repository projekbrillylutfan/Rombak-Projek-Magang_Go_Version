package domain

import "time"

type Bupati struct {
	ID             int64     `gorm:"primary_key;column:id_bupati;autoIncrement"`
	Nama           string    `gorm:"column:nama"`
	PeriodeJabatan string    `gorm:"column:periode_jabatan"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt       time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (b *Bupati) TableName() string {
	return "bupati"
}
