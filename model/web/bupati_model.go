package web

import "time"

type BupatiCreateOrUpdate struct {
	Nama           string `json:"nama" validate:"required"`
	PeriodeJabatan string `json:"periode_jabatan" validate:"required"`
}

type BupatiModel struct {
	ID             int64     `json:"id" validate:"required"`
	Nama           string    `json:"nama" validate:"required"`
	PeriodeJabatan string    `json:"periode_jabatan" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"updated_at"`
}
