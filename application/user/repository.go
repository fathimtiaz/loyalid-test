package user

import (
	"context"
	"loyalid-test/domain"
)

type Repository interface {
	GetUserByAuth0Id(context.Context, string) (domain.User, error)
}
