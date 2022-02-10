package main

import (
	"context"
	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/NetEase-Media/ngo/server"
	_ "net/http/pprof"
	"runtime"

	"github.com/ngo/ngo-demo/cmd/router"
	"github.com/ngo/ngo-demo/pkg/service"
)

func main() {
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
	s := server.Init()                 // 初始化服务器

	s.PreStart = func() error {
		log.Info("do pre-start...")
		router.InitRouter(s)   // 初始化路由
		router.InitConflict(s) // 初始化出现冲突的路由规则，只是做一个demo来使用
		service.Init()         // 初始化service中用到的工具
		return nil
	}

	s.PreStop = func(ctx context.Context) error {
		log.Info("do pre-stop...")
		return nil
	}

	s.Start() // 启动服务
}
