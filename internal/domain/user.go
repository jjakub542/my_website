package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"password_hash"`
	IsSuperuser  bool      `json:"is_superuser"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Active       bool      `json:"active"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&u.Password, validation.Required, validation.Length(3, 99)),
	)
}

func (u *User) CreatePasswordHash() {
	hashed_password := sha256.New()
	hashed_password.Write([]byte(u.Password))

	u.PasswordHash = hex.EncodeToString(hashed_password.Sum(nil))
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetOneById(string) (*User, error)
	GetOneByEmail(string) (*User, error)
	CreateOne(*User) error
	UpdateOneById(string, *User) error
	DeleteOneById(string) error
}
