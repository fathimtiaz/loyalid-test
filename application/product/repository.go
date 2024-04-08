package product

import (
	"context"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

type Repository interface {
	CreateProduct(ctx context.Context, product domain.Product) error
	GetProducts(ctx context.Context, filter repository.ProductFilter) ([]domain.Product, error)
}
