package product

import (
	"context"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

// stub for unit testing. do not use outside of test codes
type RepoStub struct {
	Repository
	Err      error
	Products []domain.Product
}

func (s RepoStub) CreateProduct(ctx context.Context, product *domain.Product) error {
	return s.Err
}

func (s RepoStub) GetProducts(ctx context.Context, filter repository.ProductFilter) ([]domain.Product, error) {
	return s.Products, s.Err
}
