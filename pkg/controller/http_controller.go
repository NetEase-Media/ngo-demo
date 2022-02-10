package controller

import (
	"encoding/json"

	"github.com/NetEase-Media/ngo/adapter/protocol"
	"github.com/gin-gonic/gin"
	"github.com/ngo/ngo-demo/pkg/service"
)

// HTTPGet Get 请求
func HTTPGet(c *gin.Context) {
	rs, err := service.HttpRequest(c.Request.Context())
	if err != nil {
		c.JSON(protocol.ErrorJsonBody(protocol.ThirdServiceError))
		return
	}
	c.JSON(protocol.JsonBody(json.RawMessage(rs)))
	return
}
