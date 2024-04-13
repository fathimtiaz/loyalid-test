package domain

import (
	"context"
	"loyalid-test/lib/jwt"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
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
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		return ""
	}

	customClaims, ok := claims.CustomClaims.(jwt.CustomClaims)
	if !ok {
		return ""
	}

	return customClaims.Nickname
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
