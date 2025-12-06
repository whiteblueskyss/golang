package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) ([]string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ListProducts(ctx context.Context) ([]string, error) {
	// implement the logic to fetch products from the data source
	return []string{"Product1", "Product2", "Product3"}, nil
}
