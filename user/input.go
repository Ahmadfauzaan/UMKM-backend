package user

type RegisterUserInput struct {
	Nama      string `json:"nama" binding:"required"`
	PeranUser string `json:"peran_user" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Telfon    string `json:"telfon" binding:"required"`
	Npwp      string `json:"npwp" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type FormCreateUserInput struct {
	Nama      string `form:"nama" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	PeranUser string `json:"peran_user" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Error     error
}

type FormUpdateUserInput struct {
	ID        int
	Nama      string `form:"nama" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	PeranUser string `json:"peran_user" binding:"required"`
	Error     error
}
