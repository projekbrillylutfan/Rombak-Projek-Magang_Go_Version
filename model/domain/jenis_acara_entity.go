package domain

import "time"

type JenisAcara struct {
	ID             int64     `gorm:"primary_key;column:id_jenis_acara;autoIncrement"`
	NamaJenisAcara string    `gorm:"column:nama_jenis_acara"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt       time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Agendas      []Agenda  `gorm:"foreignKey:IDJenisAcara" json:"agendas"`
}

func (ja *JenisAcara) TableName() string {
	return "jenis_acara"
}
