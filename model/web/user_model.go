package web

type UserCreateOrUpdate struct {
	Nama     string `json:"nama" validate:"required"`
	Jabatan  string `json:"jabatan" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserRegister struct {
	Nama     string `json:"nama" validate:"required"`
	Jabatan  string `json:"jabatan" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type UserModel struct {
	ID       int64  `json:"id" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Jabatan  string `json:"jabatan" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserEmail struct {
	Email string `json:"email" validate:"required"`
}

type RestPassUser struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}
