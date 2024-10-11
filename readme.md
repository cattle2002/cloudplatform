│  config.json  配置文件
│  go.mod
│  go.sum
│  main.go      main函数
│  readme.md
│  tree.txt
│  websocket_server_demo.go Websocket Server Demo
│  
│      
├─config
│      config.go  配置包
│      
├─core
│  │  core.go    核心组件
│  │  
│  ├─apiServer
│  │  │  
│  │  │  apiRouter.go 路由
│  │  │  apiServer.go Web服务器接口实现
│  │  │  
│  │  ├─docs
│  │  └─middleware
│  │          middleWare_Cors.go 跨域
│  │          middleWare_Jwt.go  Jwt
│  │          middleWare_Rbac.go Casbin(用户角色权限)
│  │          
│  └─othter
│          other.go 其他
│          
├─db
│      mysql.go mysql连接
│      
├─log
│      log.go 日志
│      logger.go
│      
├─logs
│      cloud.log
│      
├─upload
│      localFile.go
│      minioFile.go
│      
└─ws
│  connect.go
│  login_handler.go  登录处理
│  msg_decode.go   消息解码器
│  msg_encode.go   消息编码
│  resouce_handler.go  资源处理
│  warning_handler.go  警告处理
│  
└─protocol
protocol_ret_v1.go 协议响应
protocol_v1.go  协议请求 
            
