package product

import (
	"context"
	"loyalid-test/domain"
	"loyalid-test/repository"
)

func (s *Service) GetProducts(ctx context.Context, filter repository.ProductFilter) (products []domain.Product, err error) {
	return s.repo.GetProducts(ctx, filter)
}
