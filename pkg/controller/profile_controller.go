package controller

import (
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func Block(c *gin.Context) {
	<-time.After(time.Second)
	c.String(http.StatusOK, "ok")
}

func Mutex(c *gin.Context) {
	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
	c.String(http.StatusOK, "ok")
}

func Sys(c *gin.Context) {
	result := make(map[string]interface{})
	result["numCPU"] = runtime.NumCPU()
	result["GOMAXPROCS"] = runtime.GOMAXPROCS(0)
	c.JSON(http.StatusOK, result)
}
