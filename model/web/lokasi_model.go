package web

type LokasiCreateOrUpdate struct {
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
}

type LokasiModel struct {
	ID        int64  `json:"id" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}
