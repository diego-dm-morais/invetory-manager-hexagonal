package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_devera_validar_os_campos_obrigatorios(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	item := NewItem("Chocolate", 7.45, 20)

	isValid, _ := item.EValido()

	assert.True(isValid)
}

func Test_devera_validar_item_sem_descricao(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	item := NewItem("", 7.45, 20)

	isValid, err := item.EValido()

	assert.Falsef(isValid, err.Error())
	assert.Equal(err.Error(), "descrição do item está em branco")
}

func Test_devera_validar_item_com_preco_invalido(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	item := NewItem("Chocololate", 0, 20)
	item2 := NewItem("Refrigerante", -1.3, 20)

	isValid, err := item.EValido()
	isValid2, err2 := item2.EValido()

	assert.Falsef(isValid, err.Error())
	assert.Equal(err.Error(), "preço do item é inválido")

	assert.Falsef(isValid2, err.Error())
	assert.Errorf(err2, "preço do item é inválido")
}

func Test_devera_validar_item_com_quantida_invalida(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	item := NewItem("Chocololate", 7.5, -1)

	isValid, err := item.EValido()

	assert.Falsef(isValid, err.Error())
	assert.Errorf(err, "Quantidade do item é inválido")
}
