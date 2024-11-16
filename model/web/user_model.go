package web

type UserCreateOrUpdate struct {
	Nama     string `json:"nama" validate:"required"`
	Jabatan  string `json:"jabatan" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserModel struct {
	ID       int64  `json:"id" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Jabatan  string `json:"jabatan" validate:"required"`
	Username string `json:"username" validate:"required"`
}
