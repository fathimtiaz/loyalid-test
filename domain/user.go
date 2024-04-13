package domain

import (
	"context"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func Auth0IdCtx(ctx context.Context) string {
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		return ""
	}

	return claims.RegisteredClaims.Subject
}
