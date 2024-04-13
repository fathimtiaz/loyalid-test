package user

import (
	"context"

	"loyalid-test/domain"
)

func (s *Service) CurrentUser(ctx context.Context) (user domain.User, err error) {
	return s.repo.GetUserByUsername(ctx, domain.UsernameCtx(ctx))
}
