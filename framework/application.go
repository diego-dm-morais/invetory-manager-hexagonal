package framework

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"invetory-manager-hexagonal/framework/controller"
	"log"
	"os"
)

type Application struct {
}

func (a Application) Init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	log.Println("Variáveis foram carregadas")
	port := os.Getenv("PORT_SERVER")

	inventarioController := controller.NewInventarioController()

	e := echo.New()
	e.POST("/inventario", inventarioController.Save)
	e.Logger.Fatal(e.Start(":" + port))
}
