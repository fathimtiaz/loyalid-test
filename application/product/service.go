package product

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{}
}
