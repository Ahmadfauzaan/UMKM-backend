package user

import "time"

type User struct {
	ID             int
	Nama           string
	Pekerjaan      string
	Email          string
	Telfon         string
	Npwp           string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
