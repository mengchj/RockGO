package RockGO

import (
	"github.com/zllangct/RockGO/component"
	"github.com/zllangct/RockGO/launcher"
	"runtime"
	"time"
)

var defaultNode *Server

type Server = launcher.LauncherComponent

//新建一个服务节点
func NewServerNode() *Server {
	//构造运行时
	runtime:=Component.NewRuntime(Component.Config{ThreadPoolSize: runtime.NumCPU()})
	runtime.UpdateFrameByInterval(time.Millisecond*100)

	//构造启动器
	launcher:=&launcher.LauncherComponent{}
	runtime.Root().AddComponent(launcher)
	return launcher
}

//获取默认节点
func DefaultServer() *Server {
	if defaultNode != nil {
		return defaultNode
	}
	defaultNode = NewServerNode()
	return defaultNode
}
