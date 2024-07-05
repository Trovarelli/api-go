package main

import (
	"api-curriculos/controller"
	db "api-curriculos/internal/database"
	"api-curriculos/repository"
	usecase "api-curriculos/useCase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de Repositorio
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada usecase
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)

	//Camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductController.GetProducts)

	server.GET("/products/:productId", ProductController.GetProductById)

	server.POST("/products", ProductController.CreateProduct)

	server.Run(":8000")
}
