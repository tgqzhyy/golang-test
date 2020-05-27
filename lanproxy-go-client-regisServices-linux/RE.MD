#这是一个改写版
这是原版server
github.com/ffay/lanproxy
这是原版客户端
lanproxy-go-client https://github.com/ffay/lanproxy-go-client

# lanproxy-go-client-RegisServices for linux
介是给Linux的，代码冗余就冗余吧，没想过整合到一起，要考虑的是快速应用起来。
用屁股飞上天都可以，只要能运转，代码写的和屎一样又怎么了
都不能运转起来，简直比屎还垃圾。
```text
var conf *tls.Config
start("key", "IP", 4900, conf) // TODO 修改这里固化配置信息

需要用的库https://github.com/kardianos/service
start, stop, restart, install, uninstall 可用于操作服务
sudo fuckRegisServicesLinux   install   -- 安装服务
sudo fuckRegisServicesLinux   start     -- 启动服务
sudo fuckRegisServicesLinux   stop     -- 停止服务

service fuckRegisServicesLinux start
```
