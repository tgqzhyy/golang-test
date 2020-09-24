package router

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	TRACE   = "TRACE"
)

var GroupList = make([]*routerGroup, 0)

var PermissionMap = make(map[string]string, 0)

// 路由信息
type router struct {
	Method       string            //方法名称
	RelativePath string            //url路径
	Permiss      string            //权限字符串
	HandlerFunc  []gin.HandlerFunc //执行函数
}

// 路由组信息
type routerGroup struct {
	ServerName   string            // 服务名称
	RelativePath string            // url路径
	Handlers     []gin.HandlerFunc //中间件
	Router       []*router         // 路由信息
}

// 根据url获取权限字符串
func FindPermission(url string) string {
	return PermissionMap[url]
}

//创建一个路由组
func New(serverName, relativePath string, middleware ...gin.HandlerFunc) *routerGroup {
	var rg routerGroup
	rg.ServerName = serverName
	rg.Router = make([]*router, 0)
	rg.RelativePath = relativePath
	rg.Handlers = middleware
	GroupList = append(GroupList, &rg)
	return &rg
}

// 添加路由信息
func (group *routerGroup) Handle(method, relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	var r router
	r.Method = method
	r.Permiss = permiss
	r.RelativePath = relativePath
	r.HandlerFunc = handlers
	group.Router = append(group.Router, &r)
	if len(permiss) > 0 {
		if strings.EqualFold(relativePath, "/") { //字符串比较是否相等,忽略大小写
			PermissionMap[group.RelativePath] = permiss
		} else {
			PermissionMap[group.RelativePath+relativePath] = permiss
		}
	}
	return group
}

//添加路由信息-ANY
func (group *routerGroup) ANY(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle("ANY", relativePath, permiss, handlers...)
	return group
}

//添加路由信息-GET 请求获取由Request_URL所标识的资源
func (group *routerGroup) GET(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(GET, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-POST 在Request_URL所标识的资源后附件新的数据
func (group *routerGroup) POST(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(POST, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-OPTIONS 请求查询服务器的性能，或者查询与资源相关的选择和需求
func (group *routerGroup) OPTIONS(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(OPTIONS, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-PUT 请求服务器存储一个资源，并用Request_URL作为其标识x
func (group *routerGroup) PUT(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PUT, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-PATCH PATCH方法出现的较晚，它在2010年的RFC 5789标准中被定义。PATCH请求与PUT请求类似，同样用于资源的更新
func (group *routerGroup) PATCH(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PATCH, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-HEAD 请求获取由Request_URL所标识的资源的响应消息报头
func (group *routerGroup) HEAD(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(HEAD, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-DELETE 请求服务器删除由Request_URL所标识的资源
func (group *routerGroup) DELETE(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(DELETE, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-CONNECT CONNECT方法是HTTP/1.1协议预留的，能够将连接改为管道方式的代理服务器。通常用于SSL加密服务器的链接与非加密的HTTP代理服务器的通信
func (group *routerGroup) CONNECT(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(CONNECT, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-TRACE 请求服务器会送收到的请求消息，用于测试或诊断
func (group *routerGroup) TRACE(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(TRACE, relativePath, permiss, handlers...)
	return group
}
