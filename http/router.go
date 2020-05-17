package http

import (
	"github.com/gin-gonic/gin"
	"notesapp/notes/routes"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/notes", routes.GetNotesRoute)

	return router
}