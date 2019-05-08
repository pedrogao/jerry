# gin-rest-template

> 一个使用gin开发restful api应用的模板

## 代理

**国内容易被墙，推荐设置proxy**

```sh
export GOPROXY=https://goproxy.io
```

## 配置

使用viper读取和使用配置，环境变量的前缀的`JERRY`。


## 运行

### 懒人版

```bash
./run.sh
```

### 详细版

先进行编译：

```bash
export GOPROXY=https://goproxy.io

go build -o jerry jerry.go
```

由于使用的是go module，在build的期间会默认下载需要的所有依赖，等待时间较长，耐心等待。

运行：

```bash
export JERRY_ADDR=:5000
 
export JERRY_RUNMODE=release
 
./jerry
```

## TODO LIST

- [x] 全局异常处理
- [x] 参数检验
- [x] 多级路由，路由分层，路由前缀
- [x] JWT 支持
- [x] 日志记录
- [x] ORM 框架集成
- [x] 配置文件驱动
- [x] cors跨域
- [x] 测试驱动
- [ ] 优化，美观

> 以高度封装带来的简单性，并不会带来真正的简单，只会带来更加的复杂