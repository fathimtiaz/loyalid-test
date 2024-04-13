package user

import (
	"context"

	"loyalid-test/domain"
)

func (s *Service) CurrentUser(ctx context.Context) (user domain.User, err error) {
	return s.repo.GetUserByAuth0Id(ctx, domain.Auth0IdCtx(ctx))
}
