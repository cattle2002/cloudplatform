package main

import (
	"cloudplatform/config"
	"cloudplatform/core"
	"cloudplatform/db"
	"cloudplatform/log"
	"cloudplatform/ws"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadConfig()
	err := log.Init(config.LogConfig())
	if err != nil {
		fmt.Printf("初始化日志错误:%s\r\n", err.Error())
		os.Exit(1)
	}
	err = db.GetMysql(config.MysqlConfig())
	if err != nil {
		log.Errorf("初始化MYSQL错误:%s\r\n", err.Error())
		os.Exit(1)
	}
	ws.Connect(config.SuanliConfig())
	//fmt.Println("???")
	Core := core.New()
	err = Core.Init()
	if err != nil {
		log.Errorf("初始化核心组件错误:%s\r\n", err.Error())
		os.Exit(1)
	}

	ControlCtx, cancel := context.WithCancel(context.Background())
	Core.Run(ControlCtx)
	// 处理SIGINT和SIGTERM信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	sig := <-interrupt
	log.Info("接受中断信号:", sig)
	cancel()
	Core.Close()
	Core.Wait()
}
