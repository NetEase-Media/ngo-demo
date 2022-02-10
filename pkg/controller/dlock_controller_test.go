package controller

import (
	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/NetEase-Media/ngo/client/httplib"
	"github.com/NetEase-Media/ngo/server"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestDlock(t *testing.T) {
	path := "/dlock"
	s := server.Init()
	s.AddRoute("GET", path, Dlock)
	go s.Start()
	defer func() {
		s.Stop()
	}()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			str := strconv.Itoa(rand.Int())
			var rs string
			// 端口最好能获取
			code, err := httplib.Get("http://localhost:8080/" + path + "?str=" + str).BindString(&rs).Do()
			if err != nil {
				log.Error(err)
			}
			println(str, rs)
			assert.Equal(t, code, 200)
			assert.Equal(t, str, rs)
			wg.Done()
		}()
	}
	wg.Wait()
}
