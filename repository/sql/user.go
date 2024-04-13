package sql

import (
	"context"
	"loyalid-test/domain"
)

func (repo *SQLRepository) CreateUser(ctx context.Context, user *domain.User) (err error) {
	err = repo.db.QueryRowContext(ctx, `
		INSERT INTO user_ (id, username, created_at) 
		VALUES ($1, $2, $3) RETURNING id
		`, user.Id, user.Username, user.CreatedAt).Scan(&user.Id)
	if err != nil {
		return err
	}

	return
}

func (repo *SQLRepository) GetUserByUsername(ctx context.Context, username string) (user domain.User, err error) {
	err = repo.db.QueryRowContext(ctx, `
		SELECT id, username, created_at
		FROM user_
		WHERE username = $1
		`, username).Scan(&user.Id, &user.Username, &user.CreatedAt)
	if err != nil {
		return
	}

	return
}

func (repo *SQLRepository) GetUserByAuth0Id(ctx context.Context, auth0Id string) (user domain.User, err error) {
	err = repo.db.QueryRowContext(ctx, `
		SELECT id, username, created_at
		FROM user_
		WHERE auth0_id = $1
		`, auth0Id).Scan(&user.Id, &user.Username, &user.CreatedAt)
	if err != nil {
		return
	}

	return
}
