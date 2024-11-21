package web

import "time"

type AgendaCreateOrUpdate struct {
	NamaAgenda     string    `json:"nama_agenda" validate:"required"`
	IDBupati       int64     `json:"id_bupati" validate:"required"`
	Deskripsi      string    `json:"deskripsi" validate:"required"`
	IDLokasi       int64     `json:"id_lokasi" validate:"required"`
	IDJenisAcara   int64     `json:"id_jenis_acara" validate:"required"`
	TanggalMulai   time.Time `json:"tanggal_mulai" validate:"required"`
	TanggalSelesai time.Time `json:"tanggal_selesai" validate:"required"`
}

type AgendaModel struct {
	ID             int64           `json:"id_agenda" validate:"required"`
	NamaAgenda     string          `json:"nama_agenda" validate:"required"`
	Bupati         *BupatiModel     `json:"bupati" validate:"required"`
	Deskripsi      string          `json:"deskripsi" validate:"required"`
	Lokasi         *LokasiModel     `json:"lokasi" validate:"required"`
	JenisAcara     *JenisAcaraModel `json:"jenis_acara" validate:"required"`
	TanggalMulai   time.Time       `json:"tanggal_mulai" validate:"required"`
	TanggalSelesai time.Time       `json:"tanggal_selesai" validate:"required"`
}
