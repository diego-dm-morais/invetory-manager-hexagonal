package controller

type ItemDto struct {
	Descricao  string  `json:"description"`
	Preco      float64 `json:"price"`
	Quantidade int     `json:"amount"`
}
