package domain

import "github.com/google/uuid"

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p *Product) GenerateId() {
	p.Id = uuid.NewString()
}