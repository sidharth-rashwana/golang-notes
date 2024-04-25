package main

import (
	"gin-framework-use/controllers"
	"gin-framework-use/services"

	internal "gin-framework-use/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router
	router := gin.Default()
	db := internal.InitDB()

	if db == nil {
		// error while connecting to DB
		return
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)
	notesController := &controllers.NotesController{}
	notesController.InitController(*notesService)
	notesController.InitRoutes(router)

	authService := services.InitAuthService(db)

	authController := controllers.InitAuthController(authService)
	authController.InitRoutes(router)
	router.Run(":8000")
}
