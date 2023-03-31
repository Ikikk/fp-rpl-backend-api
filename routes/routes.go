package routes

import (
	"FP-RPL-ECommerce/controller"
	"FP-RPL-ECommerce/middleware"
	"FP-RPL-ECommerce/services"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, UserController controller.UserController, CustContoller controller.CustController, SellerContoller controller.SellerController, jwtSvc services.JWTService) {
	router := route.Group("")
	{
		router.POST("/register", UserController.Register)
<<<<<<< HEAD
		router.POST("/search", UserController.GetSellerByName)
		// router.GET("", middleware.Authenticate(""))

=======
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	}

	custRouter := route.Group("/customer")
	{
		custRouter.POST("/login", CustContoller.LoginCust)
<<<<<<< HEAD
		custRouter.PUT("/update", middleware.Authenticate(jwtSvc, "customer"), CustContoller.UpdateProfileCust)
		custRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "customer"), CustContoller.DeleteCust)
		custRouter.GET("/profile", middleware.Authenticate(jwtSvc, "customer"), CustContoller.ShowCustByID)
=======
		custRouter.PUT("/update", middleware.Authenticate(jwtSvc, "customer"), CustContoller.UpdateCust)
		custRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "customer"), CustContoller.DeleteCust)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	}

	sellerRouter := route.Group("/seller")
	{
		sellerRouter.POST("/login", SellerContoller.LoginCust)
<<<<<<< HEAD
		sellerRouter.PUT("/update", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.UpdateProfileSeller)
		sellerRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.DeleteSeller)
		sellerRouter.GET("/profile", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.ShowSellerByID)
=======
		sellerRouter.PUT("/update", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.UpdateSeller)
		sellerRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.DeleteSeller)

>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	}

	adminRouter := route.Group("/admin")
	{
<<<<<<< HEAD
		// adminRouter.GET("/customer/:id", CustContoller.GetCustByID)
		adminRouter.GET("/customer/all", CustContoller.GetAllCust)

		// adminRouter.GET("/seller/:id", SellerContoller.GetSellerByID)
=======
		adminRouter.GET("/customer/all", CustContoller.GetAllCust)

>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		adminRouter.GET("/seller/all", SellerContoller.GetAllSeller)
	}
}
