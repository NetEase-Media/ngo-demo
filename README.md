# ngo-demo

#### 简介
> 此项目中包含了`ngo`框架的一些常规操作，目录排版方式也是我们推荐的格式，请大家斟酌使用。

工程目录说明
```
   ngo-demo
    ├─cmd
    │  ├─router
    │  │  └─router.go           --路由目录
    │  └─main.go                --main
    ├─configs
    │  └─app-{env}.yaml         --配置文件 env=local|dev|test|pre|pro 等
    ├─logs                      --日志目录,忽略
    ├─pkg                       --业务目录
    │  ├─controller
    │  └─service
    ├─.gitignore
    ├─go.mod
    └─go.sum

```

#### 服务启动
go run ./cmd/main.go -c ./configs/app-local.yaml 

#### 建议环境
自己可以在本地搭建`mysql`, `redis`, `kafka`。
或者自己修改配置文件，连接其他主机的环境。
