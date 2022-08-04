package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_devera_somar_o_total_do_preco_do_pedido(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	var itens = []Item{
		NewItem("Chocolate", 10.2, 20),
		NewItem("Refrigerante", 7.4, 100),
		NewItem("Bolacha", 4.5, 70),
	}
	inventario := NewInventario("inventário 2022", itens)

	assert.Equal(22.1, inventario.Total())
}

func Test_devera_validar_dados_do_inventario(t *testing.T) {
	assert := assert.New(t)

	var itens = []Item{
		NewItem("Chocolate", 10.2, 20),
		NewItem("Refrigerante", 7.4, 100),
		NewItem("Bolacha", 4.5, 70),
	}

	inventario := NewInventario("inventário 2022", itens)
	isValid, err := inventario.EValido()

	assert.True(isValid, err)
}

func Test_devera_inventário_sem_items(t *testing.T) {
	assert := assert.New(t)

	itens := []Item{}

	inventario := NewInventario("inventário 2022", itens)

	isValid, err := inventario.EValido()

	assert.False(isValid, err)
	assert.Falsef(isValid, err.Error())
	assert.Errorf(err, "Itens não encontrados")
}

func Test_devera_retornar_erro_no_caso_de_um_item_invalido(t *testing.T) {
	assert := assert.New(t)

	itens := []Item{
		NewItem("Chocolate", 0, 20),
		NewItem("Refrigerante", 7.4, 100),
		NewItem("Bolacha", 4.5, 70),
	}

	inventario := NewInventario("inventário 2022", itens)
	isValid, err := inventario.EValido()

	assert.False(isValid, err)
	assert.Falsef(isValid, err.Error())
	assert.Errorf(err, "preço do item é inválido")
}
