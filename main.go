package main

import (
	"github.com/sidharth-rashwana/notes-app/controllers"
	"github.com/sidharth-rashwana/notes-app/services"

	internal "github.com/sidharth-rashwana/notes-app/internal/database"

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
