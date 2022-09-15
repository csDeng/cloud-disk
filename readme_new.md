# 开发指南

## 开发人员需要做什么

关注业务代码编写，将重复性、与业务无关的工作交给goctl，生成好rpc服务代码后，开发人员仅需要修改

- 服务中的配置文件编写(etc/xx.json、internal/config/config.go)
- 服务中业务逻辑编写(internal/logic/xxlogic.go)
- 服务中资源上下文的编写(internal/svc/servicecontext.go)



## grpc 单独调试

[fullstorydev/grpcui: An interactive web UI for gRPC, along the lines of postman (github.com)](https://github.com/fullstorydev/grpcui)

```shell
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

grpcui -plaintext localhost:12345
```



# 相关命令

* protobuf 生成 rpc

```shell
# 进入 pb 所在文件
goctl rpc protoc xx.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../
```

* api 生成 http文件

```shell
# 进入 api 所在目录
goctl api go -api xx.api -dir . -style go_zero
```

* sql 生成 model

```shell
# 进入 sql 文件所在目录
goctl model mysql ddl -src="./*.sql" -dir="../model" -c 
```

注： 

> 可以在 model 生成的文件夹自定义dao，重新生成就不会覆盖了



 
