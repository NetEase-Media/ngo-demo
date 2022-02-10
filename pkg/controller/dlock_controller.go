package controller

import (
	"context"
	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/NetEase-Media/ngo/dlock"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

var sharedData string

// Dlock 测试分布式锁
func Dlock(c *gin.Context) {
	str := c.DefaultQuery("str", "")
	ctx := context.Background()
	succ, executed, err := dlock.NewMutex("test", func() {
		log.Info("start working...")
		sharedData = str
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		log.Info("end working...")
	}).DoContext(ctx)
	if !succ {
		log.Errorf("dlock failed, name: %s, executed: %t, err: %+v", succ, executed, err)
	}
	c.String(http.StatusOK, sharedData)
	return
}
