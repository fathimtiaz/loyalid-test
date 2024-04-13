package product

import (
	"context"
	"errors"
	"loyalid-test/domain"
	"testing"
)

func TestService_CreateProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product *domain.Product
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		{
			s: &Service{
				RepoStub{},
			},
			args: args{
				product: &domain.Product{},
			},
		},
		{
			s: &Service{
				RepoStub{
					Err: errors.New("sql error"),
				},
			},
			args: args{
				product: &domain.Product{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CreateProduct(tt.args.ctx, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
