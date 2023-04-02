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

	productRepository := repository.NewProductRepo(db, sellerRepository)
	productService := services.NewProductSvc(productRepository)
	productController := controller.NewProductController(productService, jwtService)

	wishlistRepository := repository.NewWishlistRepo(db, productRepository)
	wishlistService := services.NewWishlistSvc(wishlistRepository)
	wishlistController := controller.NewWishlistController(wishlistService, jwtService)

	reviewRepository := repository.NewReviewRepo(db)
	reviewService := services.NewReviewSvc(reviewRepository)
	reviewController := controller.NewReviewController(reviewService)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORS())

	routes.Routes(server, userController, custController, sellerController, productController, wishlistController, reviewController, jwtService)

	port := os.Getenv("PORT")
	if port == " " {
		port = "9999"
	}
	server.Run(":" + port)
}
