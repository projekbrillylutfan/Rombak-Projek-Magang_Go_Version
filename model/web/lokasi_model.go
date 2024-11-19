package web

type LokasiCreateOrUpdate struct {
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
}
