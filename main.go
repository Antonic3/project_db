package main

import (
	"project-db/controller"
	"project-db/lib"
	"project-db/model"
	"project-db/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.InitializeDatabase()
	if err != nil {
		panic(err.Error())
	}

	ordersRepository := repository.NewOrdersRepository(db)
	ordersController := controller.NewOrdersRepository(ordersRepository)

	err = db.AutoMigrate(&model.Orders{}, &model.Items{})
	if err != nil {
		panic(err.Error())
	}

	ginEngine := gin.Default()

	ginEngine.GET("/orders", ordersController.GetAll)
	ginEngine.POST("/orders", ordersController.Create)
	ginEngine.PUT("/orders/:id", ordersController.Update)
	ginEngine.DELETE("/orders/:id", ordersController.Delete)

	err = ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err.Error())
	}
}
