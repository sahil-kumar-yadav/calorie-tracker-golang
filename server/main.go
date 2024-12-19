package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sahil-kumar-yadav/calorie-tracker-golang/routes"
)

func main() {
	port := os.Getenv("PORT") //
	if port == "" {
		port = "8000"
		// log.Fatal("$PORT must be set")
	}

	router := gin.New()       // router
	router.Use(gin.Logger())  // logging
	router.Use(cors.Defult()) // default cores

	router.Post("/entry/create", routes.AddEntry)
	router.GET("entries", router.GetEntries)
	router.GET("entry/:id", routes.EntrybyId)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	//
	router.PUT("/entry/update/:id", routes.UdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	router.Run(":" + port) // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
