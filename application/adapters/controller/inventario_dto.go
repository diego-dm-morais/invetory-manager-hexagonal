package controller

type InventarioDto struct {
	Titulo    string    `json:"title"`
	Itens     []ItemDto `json:"items"`
	IdUsuario string    `json:"user_id"`
}
