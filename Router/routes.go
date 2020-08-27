package Router

import (
	"SudokuAPI/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	grp1 := router.Group("/sudoku")
	{
		grp1.GET("generate/:level", Controllers.Generator)
		grp1.POST("validate", Controllers.Validate)
	}

	return router
}
