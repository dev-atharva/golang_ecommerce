package main

import (
	"log"

	"github.com/dev-atharva/golang_ecommerce/controllers"
	"github.com/dev-atharva/golang_ecommerce/database"
	"github.com/dev-atharva/golang_ecommerce/middleware"
	"github.com/dev-atharva/golang_ecommerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addToCart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run())
}
