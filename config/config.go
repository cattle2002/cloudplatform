package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

// Server 配置结构体(本地如果启动Web服务器)
type Server struct {
	Address   string `json:"Address"`
	KeyFile   string `json:"KeyFile"`
	CertFile  string `json:"CertFile"`
	EnableTls bool   `json:"EnableTls"`
}

// Store 配置结构体(Minio  本地存储)
type Store struct {
	Address   string `json:"Address"`
	AccessID  string `json:"AccessID"`
	AccessKey string `json:"AccessKey"`
	EnableTls bool   `json:"EnableTls"`
}

// Mysql 配置结构体(Mysql )
type Mysql struct {
	Address     string `json:"Address"`
	Scheme      string `json:"Scheme"`
	User        string `json:"User"`
	Password    string `json:"Password"`
	Name        string `json:"Name"`
	MaxPoolSize int    `json:"MaxPoolSize"`
	MaxIdleSize int    `json:"MaxIdleSize"`
}

// Log 配置结构体 (日志)
type Log struct {
	File      string `json:"File"`      //日志存放位置
	Level     string `json:"Level"`     //日志级别
	MaxAge    int    `json:"MaxAge"`    //日志最大保存时间(天)
	MaxSize   int    `json:"MaxSize"`   //单个日志文件最大大小(单位MB)
	MaxBackup int    `json:"MaxBackup"` //日志最大备份数量
}

// Suanli 配置结构体
type Suanli struct {
	Address              string `json:"Address"`           //算力平台Websocket服务器地址
	EnableTls            bool   `json:"EnableTls"`         //算力平台配置了Https 需要填写True(wss://)
	DialTimeout          int    `json:"DialTimeout"`       //连接超时时长
	HeartbeatIntervalSec int    `json:"HearthIntervalSec"` //心跳保活时间
}

// Config 包含所有配置的结构体
type Config struct {
	Server Server `json:"Server"`
	Store  Store  `json:"Store"`
	Mysql  Mysql  `json:"Mysql"`
	Log    Log    `json:"Log"`
	Suanli Suanli `json:"Suanli"`
}

var config Config
var Mutex sync.RWMutex

func ServerConfig() *Server {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return &config.Server
}

func StoreConfig() *Store {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return &config.Store
}

func MysqlConfig() *Mysql {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return &config.Mysql

}

func LogConfig() *Log {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return &config.Log
}

func SuanliConfig() *Suanli {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return &config.Suanli
}
func LoadConfig() {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("read  config  error:%s\r\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%+v\n", config)
}
