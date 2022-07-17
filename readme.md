# 网盘系统

参考： [【项目实战】基于Go-zero实现网盘系统_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1cr4y1s7H4?p=2&vd_source=620efb0bc3b2b0b7169e8564f7f527a8)

注：**改动比较大**
  
## 环境准备

* Golang 1.18
* win10 x64
* vacode
* 安装 vscode 插件
  * goctl
*  [安装 `goctl`]([Goctl 简介 | go-zero](https://go-zero.dev/cn/docs/goctl/goctl)) 

```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

> 注意要将 `goctl.exe` 所在的 `bin` 目录添加到环境变量

查看有什么 goctl 命令可以用

```shell
 D:\Github\go_cloud_disk> goctl --help
A cli tool to generate api, zrpc, model code

GitHub: https://github.com/zeromicro/go-zero
Site:   https://go-zero.dev

Usage:
  goctl [command]

Available Commands:
  api               Generate api related files
  bug               Report a bug
  completion        Generate the autocompletion script for the specified shell
  docker            Generate Dockerfile
  env               Check or edit goctl environment
  help              Help about any command
  kube              Generate kubernetes files
  migrate           Migrate from tal-tech to zeromicro
  model             Generate model code
  quickstart        quickly start a project
  rpc               Generate rpc code
  template          Template operation
  upgrade           Upgrade goctl to latest version

Flags:
  -h, --help      help for goctl
  -v, --version   version for goctl

Use "goctl [command] --help" for more information about a command.
```

看看 api 命令

```shell
PS D:\Github\go_cloud_disk> goctl api --help
Generate api related files

Usage:
  goctl api [flags]
  goctl api [command]

Available Commands:
  dart        Generate dart files for provided api in api file
  doc         Generate doc files
  format      Format api files
  go          Generate go files for provided api in yaml file
  java        Generate java files for provided api in api file
  kt          Generate kotlin code for provided api file
  new         Fast create api service
  plugin      Custom file generator
  ts          Generate ts files for provided api in api file
  validate    Validate api file

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -h, --help            help for api
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --o string        Output a sample api file
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
Use "goctl api [command] --help" for more information about a command.



PS D:\Github\go_cloud_disk> goctl api new --help
Fast create api service

Usage:
  goctl api new [flags]

Examples:
goctl api new [options] service-name

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -h, --help            help for new
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                                The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string    The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md] (default "gozero")
```



*  安装 `protoc`

```shell
goctl env check -i -f --verbose 
```

* `protoc-gen-go` / `protoc-gen-go-grpc`  安装

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```



## 简单使用

* 使用 goctl 生成 go-zero 框架代码

```shell
# 创建一个叫 core 的服务
goctl api new core
```

简单总结

> 生成了两个文件目录
>
> * etc
> * internal
>
> 三个文件
>
> * core.api
> * core.go
> * go.mod

```shell
|____go.mod
|____go.sum
|____greet.api // api接口与类型定义
|____etc // 网关层配置文件
| |____greet-api.yaml
|____internal
| |____config // 配置-对应etc下配置文件
| | |____config.go
| |____handler // 视图函数层, 路由与处理器
| | |____routes.go
| | |____greethandler.go
| |____logic // 逻辑处理
| | |____greetlogic.go
| |____svc // 依赖资源, 封装 rpc 对象的地方
| | |____servicecontext.go
| |____types // 中间类型
| | |____types.go
|____greet.go // main.go 入口

```



* 依赖整理

```shell
 D:\Github\go_cloud_disk\core> go mod tidy
go: finding module for package github.com/zeromicro/go-zero/core/logx
go: finding module for package github.com/zeromicro/go-zero/core/conf
... 省略
go: downloading golang.org/x/net v0.0.0-20220531201128-c960675eff93
```

* 启动服务

```shell
PS D:\Github\go_cloud_disk\core> go run core.go -f etc/core-api.yaml
Starting server at 0.0.0.0:8888...
```

> -f 是使用flag包指定的配置文件参数





* 注意，后续修改了 api 文件，都可以直接使用命令快速生成模板

```shell
goctl api go -api core.api -dir . -style go_zero
```

生成模板之后，再到 `logic` 目录去相应的文件里面写业务逻辑。



## 项目设计

数据库设计与业务逻辑设计

### 数据库

* **用户信息**：存储用户基本信息，用于登录
* **公共文件存储池**：存储文件信息
* **用户存储池**：对公共文件存储池中文件信息的引用
* **文件分享**



