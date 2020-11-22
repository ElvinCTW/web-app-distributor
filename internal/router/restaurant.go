package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-app-distributor/internal/controller"
)

func Restaurant(r *gin.Engine) {
	r.GET("/app/v1/restaurant/:id", func(c *gin.Context) {
		id := c.Param("id")
		if r := controller.ShowRestaurant(id); r.Data == nil {
			c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
		} else {
			c.JSON(http.StatusOK, r)
		}
	})
}
