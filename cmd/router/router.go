package router

import (
	"fmt"
	"github.com/NetEase-Media/ngo/server"
	"github.com/NetEase-Media/ngo/server/timeout"
	"net/http"
	"time"

	"github.com/ngo/ngo-demo/pkg/controller"
)

// InitRouter 项目路由统一管理
func InitRouter(s *server.Server) {
	//	s.AddRoute(server.GET, "/hello", controller.Hello)
	s.AddRouteWithMethods([]server.Method{server.GET, server.POST}, "/hello", controller.Hello)

	s.AddRoute(server.GET, "/hello/:somebody", controller.HelloSomebody)

	s.AddRoute(server.GET, "/say/no", controller.SayNo)

	s.AddRoute(server.GET, "/http", controller.HTTPGet)

	s.AddRouteWithMethods([]server.Method{server.GET, server.POST}, "/kafka/send", controller.Send)

	s.AddRoute(server.GET, "/mysql", controller.CarInfo)

	s.AddRoute(server.POST, "/redis/write", controller.Write)
	s.AddRoute(server.GET, "/redis/read", controller.Read)
	s.AddRoute(server.GET, "/redis/readAll", controller.ReadAll)
	s.AddRoute(server.POST, "/redis/sharded/write", controller.ShardedWrite)
	s.AddRoute(server.GET, "/redis/sharded/read", controller.ShardedRead)
	s.AddRoute(server.GET, "/redis/sharded/pipe", controller.ShardedPipe)

	s.AddRoute(server.GET, "/dlock", controller.Dlock)

	s.AddRoute(server.GET, "/block", controller.Block)
	s.AddRoute(server.GET, "/mutex", controller.Mutex)

	s.AddRoute(server.GET, "/sys", controller.Sys)

	s.AddRoute(server.GET, "/timeout", timeout.Timeout(
		timeout.WithTimeout(2*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // optional
		timeout.WithDefaultMsg("timeout"),                    // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		})), controller.Timeout)

	//s.AddRoute(server.GET, "/stop", func(c *gin.Context) {
	//	ctx, cancel := context.WithTimeout(c, time.Second)
	//	defer cancel()
	//	go s.Stop(ctx)
	//	c.String(http.StatusOK, "ok")
	//})
}
