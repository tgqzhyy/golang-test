#这是一个改写版
这是原版server
github.com/ffay/lanproxy
这是原版客户端
lanproxy-go-client https://github.com/ffay/lanproxy-go-client

# TODO
在Windows上面把客户端注册为服务，emmm有源码直接改写就好了，不使用其他外挂式方式。

# 交叉编译
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
# 注册为Windows服务
$ lanproxy-go-client-regisServices.exe -service-install
# 启动和停止Windows服务
$ lanproxy-go-client-regisServices.exe -service-start
$ lanproxy-go-client-regisServices.exe -service-stop
# 删除服务
# 删除之前需要先停止服务
$ lanproxy-go-client-regisServices.exe -service-remove
