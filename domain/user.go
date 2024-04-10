package domain

import (
	"context"
	"fmt"
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

func UsernameCtx(ctx context.Context) string {
	username, ok := ctx.Value("username").(string)
	if !ok {
		fmt.Printf("failed getting username from context: %+v", ctx.Value("username"))
	}

	return username
}

func (u *User) GenerateId() {
	u.Id = uuid.NewString()
}

func (u *User) HashPassword() (err error) {
	var hashByte []byte

	if hashByte, err = bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost); err != nil {
		return err
	}

	u.password = string(hashByte)

	return
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Password() string {
	return u.password
}
