package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-app-distributor/internal/controller"
)

func Article(r *gin.Engine) {
	r.GET(appPrefix+"/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		if a := controller.GetArticle(id); a == "" {
			c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
			return
		} else {
			c.String(http.StatusOK, a)
		}
	})
}
