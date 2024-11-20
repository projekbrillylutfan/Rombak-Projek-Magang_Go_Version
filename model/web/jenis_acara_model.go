package web

type JenisAcaraCreateOrUpdate struct {
	NamaJenisAcara string `json:"nama_jenis_acara" validate:"required"`
}

type JenisAcaraModel struct {
	ID             int64  `json:"id" validate:"required"`
	NamaJenisAcara string `json:"nama_jenis_acara" validate:"required"`
	CreatedAt      string `json:"created_at"`
	UpdateAt       string `json:"updated_at"`
}
