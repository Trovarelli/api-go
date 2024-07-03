package model

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"descricao"`
	Price float64 `json:"preco"`
}
