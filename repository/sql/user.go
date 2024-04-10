package sql

import (
	"context"
	"loyalid-test/domain"
)

func (repo *SQLRepository) CreateUser(ctx context.Context, user *domain.User) (err error) {
	err = repo.db.QueryRowContext(ctx, `
		INSERT INTO user_ (id, username, password, created_at) 
		VALUES ($1, $2, $3, $4) RETURNING id
		`, user.Id, user.Username, user.Password(), user.CreatedAt).Scan(&user.Id)
	if err != nil {
		return err
	}

	return
}

func (repo *SQLRepository) GetUserByUsername(ctx context.Context, username string) (user domain.User, err error) {
	var password string

	err = repo.db.QueryRowContext(ctx, `
		SELECT id, username, password, created_at
		FROM user_
		WHERE username = $1
		`, username).Scan(&user.Id, &user.Username, &password, &user.CreatedAt)
	if err != nil {
		return
	}

	user.SetPassword(password)

	return
}
