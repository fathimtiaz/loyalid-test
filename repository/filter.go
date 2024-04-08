package repository

type ProductFilter struct {
	Pagination
}

type Pagination struct {
	Page  int
	Limit int
}
