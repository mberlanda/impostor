package main

import (
	"github.com/mberlanda/impostor/api/controller"
	"github.com/mberlanda/impostor/api/db"
	"github.com/mberlanda/impostor/api/repository"
	"github.com/mberlanda/impostor/api/routes"
	"github.com/mberlanda/impostor/api/service"
)

func main() {

	router := routes.NewGinRouter()

	userDb := db.InitUserDb()
	userRepository := repository.NewUserRepository(&userDb)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)
	userRoute := routes.NewUserRoute(&userController, &router)
	userRoute.Setup()

	router.Gin.Run(":8080")
}
