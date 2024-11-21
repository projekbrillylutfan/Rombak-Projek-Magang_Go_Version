package domain

import "time"

type Agenda struct {
	IDAgenda       int64      `gorm:"primaryKey;autoIncrement" json:"id_agenda"`
	IDBupati       int64      `gorm:"column:id_bupati"`
	NamaAgenda     string    `gorm:"column:nama_agenda"`
	Deskripsi      string    `gorm:"column:deskripsi"`
	IDLokasi       int64      `gorm:"column:id_lokasi"`
	IDJenisAcara   int64      `gorm:"column:id_jenis_acara"`
	TanggalMulai   time.Time `gorm:"column:tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"column:tanggal_selesai"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Bupati        Bupati        `gorm:"foreignKey:IDBupati" json:"bupati"`
    Lokasi        Lokasi        `gorm:"foreignKey:IDLokasi" json:"lokasi"`
    JenisAcara    JenisAcara    `gorm:"foreignKey:IDJenisAcara" json:"jenis_acara"`
}

func (a *Agenda) TableName() string {
	return "agenda"
}
