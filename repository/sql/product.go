package sql

import (
	"context"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

func (repo *SQLRepository) CreateProduct(ctx context.Context, product *domain.Product) (err error) {
	err = repo.db.QueryRowContext(ctx, `
		INSERT INTO product_ (name_, price) 
		VALUES ($1, $2) RETURNING id
		`, product.Name, product.Price).Scan(&product.Id)
	if err != nil {
		return err
	}

	return
}

func (repo *SQLRepository) GetProducts(ctx context.Context, filter repository.ProductFilter) (products []domain.Product, err error) {
	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Page <= 0 {
		filter.Page = 0
	}

	rows, err := repo.db.QueryContext(ctx, `
		SELECT id, name_, price
		FROM product_
		LIMIT $1 OFFSET $2
		`, filter.Limit, filter.Page)
	if err != nil {
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
