package entities

import (
	"errors"
	"time"
)

type Inventario struct {
	Titulo        string
	DataDeCriacao time.Time
	Itens         []Item
}

func (inventario Inventario) EValido() (bool, error) {

	isValid := true
	var err error = nil

	if len(inventario.Itens) == 0 {
		isValid = false
		err = errors.New("Itens não encontrados")
		return isValid, err
	}

	for _, item := range inventario.Itens {
		isValid, err = item.EValido()
		if isValid == false {
			return isValid, err
		}
	}
	return isValid, err
}

func (inventario Inventario) Total() float64 {
	var valorTotal = 0.0
	for _, item := range inventario.Itens {
		valorTotal += item.Preco
	}
	return valorTotal
}

func NewInventario(tituto string, itens []Item) Inventario {
	return Inventario{
		Titulo:        tituto,
		Itens:         itens,
		DataDeCriacao: time.Now(),
	}
}
