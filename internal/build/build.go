package build

import (
	"compoze_rest_api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func BuildBookResource(router *gin.Engine, prefix string) {
	router.GET(prefix+"/:id", handlers.GetBookById)
	router.POST(prefix, handlers.CreateBook)
	router.PUT(prefix+"/:id", handlers.UpdateBook)
	router.DELETE(prefix+"/:id", handlers.DeleteBook)
}
func BuildManyBooksResourcePrefix(router *gin.Engine, prefix string) {
	router.GET(prefix, handlers.GetAllBooks)
}
