package repository

type ProductFilter struct {
	Pagination
}

type Pagination struct {
	page  int
	limit int
}

func NewPagination(page, limit int) Pagination {
	if limit <= 0 {
		limit = 10
	}

	if page <= 0 {
		page = 1
	}

	return Pagination{
		page:  page,
		limit: limit,
	}
}

func (p *Pagination) Page() int {
	return p.page
}

func (p *Pagination) Limit() int {
	return p.limit
}
