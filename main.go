package main

import (
	"fmt"
	"os"

	"FP-RPL-ECommerce/config"
	"FP-RPL-ECommerce/controller"
	"FP-RPL-ECommerce/middleware"
	"FP-RPL-ECommerce/repository"
	"FP-RPL-ECommerce/routes"
	"FP-RPL-ECommerce/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed load environment file")
	} else {
		fmt.Println("success environment file")
	}

	db := config.ConnectDB()
	jwtService := services.NewJWTService()

	custRepository := repository.NewCustRepo(db)
	custService := services.NewCustSvc(custRepository)
	custController := controller.NewCustController(custService, jwtService)

	sellerRepository := repository.NewSellerRepo(db)
	sellerService := services.NewSellerSvc(sellerRepository)
	sellerController := controller.NewSellerController(sellerService, jwtService)

	userController := controller.NewUserController(custService, sellerService)

	// defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORS())

	routes.Routes(server, userController, custController, sellerController, jwtService)

	port := os.Getenv("PORT")
	if port == " " {
		port = "9999"
	}
	server.Run(":" + port)
}
