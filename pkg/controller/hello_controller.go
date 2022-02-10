package controller

import (
	"fmt"
	"github.com/NetEase-Media/ngo/adapter/protocol"
	"github.com/gin-gonic/gin"
)

// Hello 向世界问好
func Hello(c *gin.Context) {
	c.Set("data", "data...")
	c.JSON(protocol.JsonBody("Hello, World!"))
}

// HelloSomebody
func HelloSomebody(c *gin.Context) {
	somebody := c.Param("somebody")
	c.JSON(protocol.JsonBody(fmt.Sprintf("Hello, %s!", somebody)))
}

// SayNo This World Is Not Perfect As You Wish！
func SayNo(c *gin.Context) {
	c.JSON(protocol.ErrorJsonBody(protocol.ParamsNotValid))
}
