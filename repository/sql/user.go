package sql

import (
	"context"
	"loyalid-test/domain"
)

func (repo *SQLRepository) CreateUser(ctx context.Context, user *domain.User) (err error) {
	err = repo.db.QueryRowContext(ctx, `
		INSERT INTO user_ (username, created_at) 
		VALUES ($1, $2) RETURNING id
		`, user.Username, user.CreatedAt).Scan(&user.Id)
	if err != nil {
		return err
	}

	return
}
