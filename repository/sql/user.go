package sql

import (
	"context"
	"database/sql"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

func (repo *SQLRepository) GetUserByAuth0Id(ctx context.Context, auth0Id string) (user domain.User, err error) {
	err = repo.db.QueryRowContext(ctx, `
		SELECT id, username, created_at
		FROM user_
		WHERE auth0_id = $1
		`, auth0Id).Scan(&user.Id, &user.Username, &user.CreatedAt)

	if err == sql.ErrNoRows {
		err = repository.ErrNotFound
	}

	return
}
