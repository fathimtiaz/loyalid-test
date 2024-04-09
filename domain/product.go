package domain

type Product struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
