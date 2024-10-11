# 项目结构

以下是项目的主要目录和文件结构：

- `config.json` - 配置文件。
- `go.mod` - Go模块依赖。
- `go.sum` - Go模块依赖的校验和。
- `main.go` - 包含程序的`main`函数。
- `readme.md` - 项目说明文档。
- `tree.txt` - 目录结构文本文件。
- `websocket_server_demo.go` - WebSocket服务器演示。

## 目录结构

### config
包含配置相关的Go代码。

- `config.go` - 配置包。

### core
包含项目核心组件。

- `core.go` - 核心组件代码。
- `othter`/`other.go` - 其他组件代码。

#### apiServer
包含API服务器相关的代码。

- `apiRouter.go` - API路由配置。
- `apiServer.go` - Web服务器接口实现。

##### docs
（此处未说明具体内容，可能为API文档相关）

##### middleware
包含中间件代码。

- `middleWare_Cors.go` - 处理跨域的中间件。
- `middleWare_Jwt.go` - JWT认证中间件。
- `middleWare_Rbac.go` - 基于Casbin的用户角色权限中间件。

### db
数据库相关代码。

- `mysql.go` - MySQL连接配置。

### log
日志相关代码。

- `log.go` - 日志功能实现。
- `logger.go` - 日志记录器配置。

### logs
存放日志文件。

- `cloud.log` - 云服务日志文件。

### upload
文件上传相关代码。

- `localFile.go` - 本地文件上传处理。
- `minioFile.go` - MinIO对象存储服务文件上传处理。

### ws
WebSocket相关代码。

- `connect.go` - WebSocket连接处理。
- `login_handler.go` - 登录处理。
- `msg_decode.go` - 消息解码器。
- `msg_encode.go` - 消息编码器。
- `resouce_handler.go` - 资源处理。
- `warning_handler.go` - 警告处理。

### protocol
协议相关代码。

- `protocol_ret_v1.go` - 协议响应V1。
- `protocol_v1.go` - 协议请求V1。