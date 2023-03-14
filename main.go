package main

import (
	"authgo/model/controller"
	"authgo/model/database"
	"authgo/model/middleware"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
			secured.GET("/allproducts", controller.GetProducts) //ger prodcuts
			secured.GET("/getproduct", controller.GetProduct)
			secured.GET("/addproduct", controller.AddProduct)
			secured.GET("/updateproduct", controller.UpdateProduct)
			secured.GET("/deleteproduct", controller.DeleteProduct)
		}
	}
	return router
}

func main() {

	LoadAppConfig()
	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	router := initRouter()
	router.Run(":8081")
}
