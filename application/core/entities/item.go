package entities

import (
	"errors"
)

type Item struct {
	Descricao  string
	Preco      float64
	Quantidade int
}

func (i Item) EValido() (bool, error) {
	if len(i.Descricao) == 0 {
		return false, errors.New("descrição do item está em branco")
	}
	if i.Preco <= 0 {
		return false, errors.New("preço do item é inválido")
	}
	if i.Quantidade < 0 {
		return false, errors.New("Quantidade do item é inválido")
	}
	return true, nil
}

func NewItem(descricao string, preco float64, quantidade int) Item {
	return Item{
		Preco:      preco,
		Descricao:  descricao,
		Quantidade: quantidade,
	}
}
