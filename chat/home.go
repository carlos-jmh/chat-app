package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "index.html")
}
