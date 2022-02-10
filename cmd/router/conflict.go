// @Copy Right netease.com
// author: tanxiangya
// desc:
// httprouter 本身的路由查找和构建逻辑，导致带路径参数的路由之间会存在冲突。
// 官方目前已经发布了mr，将该问题进行了修复，但是具体合并到gin的时间不确定，
// 因此，先使用下述的方式来解决路由冲突的问题

package router

import (
	"github.com/NetEase-Media/ngo/server"
	"github.com/ngo/ngo-demo/pkg/controller"
)

func InitConflict(s *server.Server) {
	// 目前这样的处理方式启动的时候会出现问题
	// error: 'v2' in new path '/nc/video/detail/v2/:vid.html' conflicts with existing wildcard ':vid.html' in existing prefix '/nc/video/detail/:vid.html'
	// solution: 目前官方以及给释放修正代码，gin合并代码
	// s.AddRoute(server.GET, "/nc/video/detail/:vid.html", controller.ActionOne)
	// s.AddRoute(server.GET, "/nc/video/detail/v2/:vid.html", controller.ActionTwo)

	// 目前解决方案
	s.AddRoute(server.GET, "/nc/video/detail/*action", controller.ActionCombine)
}
