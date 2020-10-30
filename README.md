#Grpc基础调用示例

## 安装工具

```go
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
```

生成protobuf对应的go文件

```go
cd proto
protoc --go_out=plugins=grpc:. helloworld.proto
```

plugins=grpc 是启用grpc插件生成grpc相关的代码，也就是protoc文件里service里定义的接口
