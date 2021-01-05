package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	PeranUser string `json:"peran_user"`
	Email     string `json:"email"`
	Telfon    string `json:"telfon"`
	Npwp      string `json:"npwp"`
	Token     string `json:"token"`
	ImageURL  string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Nama:      user.Nama,
		PeranUser: user.PeranUser,
		Email:     user.Email,
		Telfon:    user.Telfon,
		Npwp:      user.Npwp,
		Token:     token,
		ImageURL:  user.AvatarFileName,
	}

	return formatter
}
