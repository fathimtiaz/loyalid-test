package user

import (
	"context"
	"loyalid-test/domain"
)

type Repository interface {
	CreateUser(context.Context, *domain.User) error
	GetUserByUsername(context.Context, string) (domain.User, error)
	GetUserByAuth0Id(context.Context, string) (domain.User, error)
}
