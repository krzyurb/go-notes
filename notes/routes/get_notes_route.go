package routes

import "github.com/gin-gonic/gin"

func GetNotesRoute(c *gin.Context) {
	c.JSON(200, gin.H{"notes": "list of notes"})
}
