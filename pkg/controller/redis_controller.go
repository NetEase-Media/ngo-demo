package controller

import (
	"github.com/NetEase-Media/ngo/adapter/protocol"
	"github.com/gin-gonic/gin"
	"github.com/ngo/ngo-demo/pkg/service"
)

// Read 读取缓存
func Read(c *gin.Context) {
	rs := service.Read()
	c.JSON(protocol.JsonBody(rs))
	return
}

// ReadAll 批量读取缓存
func ReadAll(c *gin.Context) {
	rs := service.MultiRead()
	c.JSON(protocol.JsonBody(rs))
	return
}

// Write 写缓存
func Write(c *gin.Context) {
	msg, ok := c.GetQuery("msg")
	if !ok {
		c.JSON(protocol.ErrorJsonBody(protocol.ParamsNotValid))
		return
	}
	err := service.Write(msg)
	if err != nil {
		c.JSON(protocol.ErrorJsonBody(protocol.ThirdServiceError))
		return
	}
	c.JSON(protocol.JsonBody("写入成功"))
}

// ShardedWrite 写缓存
func ShardedWrite(c *gin.Context) {
	msg, ok := c.GetQuery("msg")
	if !ok {
		c.JSON(protocol.ErrorJsonBody(protocol.ParamsNotValid))
		return
	}
	err := service.ShardedWrite(msg)
	if err != nil {
		c.JSON(protocol.ErrorJsonBody(protocol.ThirdServiceError))
		return
	}
	c.JSON(protocol.JsonBody("写入成功"))
}

// ShardedRead 读取缓存
func ShardedRead(c *gin.Context) {
	rs := service.ShardedRead()
	c.JSON(protocol.JsonBody(rs))
	return
}

// ShardedPipe 读取缓存
func ShardedPipe(c *gin.Context) {
	rs := service.ShardedPipe()
	c.JSON(protocol.JsonBody(rs))
	return
}
