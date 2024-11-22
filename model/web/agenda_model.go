package web

import (
	"time"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

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

func ConvertBupatiToModel(bupati *domain.Bupati) *BupatiModel {
	return &BupatiModel{
		ID:       bupati.ID,
		Nama:     bupati.Nama,
		PeriodeJabatan:  bupati.PeriodeJabatan,
		CreatedAt: bupati.CreatedAt,
		UpdateAt: bupati.UpdateAt,
		// Tambahkan field lainnya jika diperlukan
	}
}

func ConvertLokasiToModel(lokasi *domain.Lokasi) *LokasiModel {
	return &LokasiModel{
		ID: lokasi.ID,
		Nama: lokasi.Nama,
		Alamat: lokasi.Alamat,
		CreatedAt: lokasi.CreatedAt.String(),
		UpdateAt: lokasi.UpdateAt.String(),
	}
}

func ConvertJenisAcaraToModel(jenisAcara *domain.JenisAcara) *JenisAcaraModel {
	return &JenisAcaraModel{
		ID: jenisAcara.ID,
		NamaJenisAcara: jenisAcara.NamaJenisAcara,
		CreatedAt: jenisAcara.CreatedAt.String(),
		UpdateAt: jenisAcara.UpdateAt.String(),
	}
}
