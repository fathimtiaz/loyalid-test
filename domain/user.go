package domain

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	password  string
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) GenerateId() {
	u.Id = uuid.NewString()
}

func (u *User) HashAndSetPassword(password string) (err error) {
	var hashByte []byte

	if hashByte, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return err
	}

	u.password = string(hashByte)

	return
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Password() string {
	return u.password
}
