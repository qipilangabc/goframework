# goframework
这是一个简单的go http框架实现demo，参考gin框架的实现原理
实现功能，比较简单，实现了基础的http路由转发，上下文统一封装调用。

启动服务
go run main.go
然后通过浏览器访问http://127.0.0.1:9988/member/lll/33  
浏览器输出json
{"name":"lll","age":33}

