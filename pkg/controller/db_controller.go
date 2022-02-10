package controller

import (
	"github.com/NetEase-Media/ngo/adapter/protocol"
	"github.com/gin-gonic/gin"
	"github.com/ngo/ngo-demo/pkg/service"
	"strconv"
)

//CarInfo 详情
func CarInfo(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok || idStr == "" {
		c.JSON(protocol.ErrorJsonBody(protocol.ParamsNotValid))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(protocol.ErrorJsonBody(protocol.ParamsNotValid))
		return
	}
	from, ok := c.GetQuery("from")
	t := service.GetDataFromDB(id, from)
	c.JSON(protocol.JsonBody(t))
	return
}
