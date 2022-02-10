package controller

import (
	"github.com/NetEase-Media/ngo/adapter/protocol"
	"strings"

	"github.com/gin-gonic/gin"
)

func ActionOne(c *gin.Context) {
	c.JSON(protocol.JsonBody("Hello, World!"))
}

func ActionTwo(c *gin.Context) {
	c.JSON(protocol.JsonBody("Hello, China!"))
}

func ActionCombine(c *gin.Context) {
	path := c.Param("action")
	if strings.Contains(path, "/v2") {
		ActionTwo(c)
	} else {
		ActionOne(c)
	}
	// c.JSON(protocol.JsonBody("Hello, World/China"))
}
