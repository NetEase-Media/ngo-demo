package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/gin-gonic/gin"
)

// Timeout
func Timeout(c *gin.Context) {
	log.Info("aaa")
	q, _ := c.GetQuery("t")
	t, _ := strconv.ParseInt(q, 10, 64)
	time.Sleep(time.Millisecond * time.Duration(t))
	c.String(http.StatusOK, "ok")
}
