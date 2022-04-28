package handlers

import (
	"UPDATE_WH/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(server *gin.Engine, service service.IUpdateService) {
	server.POST("/update", func(c *gin.Context) {
		process, err := strconv.ParseBool(c.GetHeader("process"))
		if err != nil {
			process = false
		}

		ctx := c.Request.Context()
		c.JSON(200, service.UpdateWh(ctx, process))
	})
}
