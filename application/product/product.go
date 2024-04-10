package product

import (
	"context"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

func (s *Service) CreateProduct(ctx context.Context, product *domain.Product) (err error) {
	product.GenerateId()

	return s.repo.CreateProduct(ctx, product)
}

func (s *Service) GetProducts(ctx context.Context, filter repository.ProductFilter) (products []domain.Product, err error) {
	return s.repo.GetProducts(ctx, filter)
}
