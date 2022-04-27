package handlers

import (
	"UPDATE_WH/service"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(server *gin.Engine, service service.IUpdateService) {
	server.POST("/update", func(c *gin.Context) {
		ctx := c.Request.Context()
		c.JSON(200, service.UpdateWh(ctx))
	})
}
