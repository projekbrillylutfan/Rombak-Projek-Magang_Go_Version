package web

type BupatiCreateOrUpdate struct {
	Nama string `json:"nama" validate:"required"`
	PeriodeJabatan string `json:"periode_jabatan" validate:"required"`
}