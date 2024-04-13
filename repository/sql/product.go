package sql

import (
	"context"
	"database/sql"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

func (repo *SQLRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return repo.db.QueryRowContext(ctx, `
		INSERT INTO product_ (id, name_, price) 
		VALUES ($1, $2, $3) RETURNING id
		`, product.Id, product.Name, product.Price).Scan(&product.Id)
}

func (repo *SQLRepository) GetProducts(ctx context.Context, filter repository.ProductFilter) (products []domain.Product, err error) {
	rows, err := repo.db.QueryContext(ctx, `
		SELECT id, name_, price
		FROM product_
		LIMIT $1 OFFSET $2
		`, filter.Limit(), (filter.Page()-1)*filter.Limit())
	if err != nil {
		if err == sql.ErrNoRows {
			err = repository.ErrNotFound
		}

		return []domain.Product{}, err
	}

	for rows.Next() {
		var product domain.Product

		if err = rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return []domain.Product{}, err
		}

		products = append(products, product)
	}

	return
}
