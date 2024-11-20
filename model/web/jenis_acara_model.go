package web

type JenisAcaraCreateOrUpdate struct {
	NamaJenisAcara string `json:"nama_jenis_acara" validate:"required"`
}
