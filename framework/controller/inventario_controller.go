package controller

import (
	"github.com/labstack/echo/v4"
	"invetory-manager-hexagonal/application/adapters/api"
	adapter_controller "invetory-manager-hexagonal/application/adapters/controller"
	"invetory-manager-hexagonal/application/adapters/repository"
	"invetory-manager-hexagonal/application/core/usecase"
	"invetory-manager-hexagonal/framework/datasource"
	framework_repository "invetory-manager-hexagonal/framework/repository"
	"net/http"
)

type InventarioController struct {
	inventarioAdapterController adapter_controller.InventarioAdapterController
}

func (i InventarioController) Save(c echo.Context) error {
	inventarioDto := new(adapter_controller.InventarioDto)
	err := c.Bind(inventarioDto)
	if err != nil {
		return err
	}

	err = i.inventarioAdapterController.Salvar(*inventarioDto)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func NewInventarioController() InventarioController {
	mongoDataSource := datasource.NewConnectorMongoDataSource()
	invetarioRepositoryMongo := framework_repository.NewInvetarioRepositoryMongo(mongoDataSource)
	inventarioAdapte := repository.NewInventarioAdapterRepository(invetarioRepositoryMongo)
	permissaoAdapterApi := api.NewPermissaoAdapterApi()
	inventarioApplication := usecase.NewInventarioApplication(inventarioAdapte, permissaoAdapterApi)
	inventarioAdapterController := adapter_controller.NewInventarioAdapterController(inventarioApplication)
	return InventarioController{
		inventarioAdapterController: *inventarioAdapterController,
	}
}
