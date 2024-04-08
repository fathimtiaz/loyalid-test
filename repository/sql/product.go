package sql

import (
	"context"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

func (repo *SQLRepository) CreateProduct(ctx context.Context, product domain.Product) (err error) {
	return
}

func (repo *SQLRepository) GetProducts(ctx context.Context, filter repository.ProductFilter) (products []domain.Product, err error) {
	return
}
