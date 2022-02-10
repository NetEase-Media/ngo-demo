package controller

import (
	"github.com/NetEase-Media/ngo/adapter/protocol"
	"github.com/gin-gonic/gin"
	"github.com/ngo/ngo-demo/pkg/service"
)

//Send 发送消息
func Send(c *gin.Context) {
	message, ok := c.GetQuery("message")
	sync, ok := c.GetQuery("sync")
	if ok {
		service.Product(message, sync)
		c.JSON(protocol.JsonBody("success"))
		return
	} else {
		c.JSON(protocol.ErrorJsonBody(protocol.ThirdServiceError))
		return
	}

}
