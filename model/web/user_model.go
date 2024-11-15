package web

type UserCreate struct {
	Nama     string `json:"nama" validate:"required"`
	Jabatan  string `json:"jabatan" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
