package controller

import (
	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/NetEase-Media/ngo/client/httplib"
	"github.com/NetEase-Media/ngo/server"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// go test -run  ^TestHello$ -args -c xxx
func TestHello(t *testing.T) {
	path := "/hello"
	s := server.Init()
	s.AddRoute("GET", path, Hello)
	go s.Start()
	defer func() {
		s.Stop()
	}()

	time.Sleep(time.Second)
	var rs string
	// 端口最好能获取
	code, err := httplib.Get("http://localhost:8080/" + path).BindString(&rs).Do()
	if err != nil {
		log.Error(err)
	}
	assert.Equal(t, 200, code)
}
