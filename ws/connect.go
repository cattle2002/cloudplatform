package ws

import (
	"bytes"
	"cloudplatform/config"
	"cloudplatform/log"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var safeWsConn *SafeWsConn
var gloConn *websocket.Conn = nil

type SafeWsConn struct {
	Mux sync.RWMutex
	//conn              *websocket.Conn
	HearthIntervalSec int
	WatcherCh         chan struct{} //断开信号重连
	Encode            func(msg []byte) ([]byte, error)
	Decode            func(msg []byte) ([]byte, error)

	Once sync.Once
}

func NewSafeWsConn(heartIntervalSec int) {
	safeWsConn = &SafeWsConn{
		//conn:              conn,
		HearthIntervalSec: heartIntervalSec,
		Encode: func(msg []byte) ([]byte, error) {
			return msg, nil
		},
		Decode: func(msg []byte) ([]byte, error) {
			return msg, nil
		},
		WatcherCh: make(chan struct{}, 10),
	}
}
func Connect(conf *config.Suanli) {
	var err error
	prefix := "ws"
	if conf.EnableTls {
		prefix = "wss"
	}
	log.Debugf("算力服务器地址:%s:%s\r\n", prefix, conf.Address)
	timeOutCtx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.DialTimeout))
	gloConn, _, err = websocket.DefaultDialer.DialContext(timeOutCtx, fmt.Sprintf("%s://%s", prefix, conf.Address), nil)
	if err != nil {
		log.Errorf("连接算力服务器错误:%s\r\n", err.Error())
		//return
	} else {
	}
	//log.Debug("connect suanli server success")
	NewSafeWsConn(conf.HeartbeatIntervalSec)
	go safeWsConn.Once.Do(func() {
		safeWsConn.Watcher(conf)
	})

	go safeWsConn.KeepLive()
	go safeWsConn.Handler()
}

//var HearthMsg = []byte("9ijnBHU*@123")

func KeepLiveMsg() []byte {
	return []byte(hex.EncodeToString([]byte("9ijnBHU*@123")))
}

func (safeWsConn *SafeWsConn) SendMsg(msg []byte) {
	safeWsConn.Mux.Lock()
	defer safeWsConn.Mux.Unlock()
	encodeMsg, _ := safeWsConn.Encode(msg)
	defer func() {
		if err := recover(); err != nil {
			log.Error("捕获业务恐慌,正在重连算力平台")
			safeWsConn.WatcherCh <- struct{}{}
		}
	}()
	if err := gloConn.WriteMessage(websocket.TextMessage, encodeMsg); err != nil {
		log.Errorf("发送消息失败:%s\r\n", err.Error())
		safeWsConn.WatcherCh <- struct{}{}
		return
	}
}
func (safeWsConn *SafeWsConn) Watcher(conf *config.Suanli) {
	for {
		var err error
		select {
		case <-safeWsConn.WatcherCh:

			log.Info("正在重连算力服务器...")

			prefix := "ws"
			if conf.EnableTls {
				prefix = "wss"
			}
			log.Debugf("算力服务器地址:%s:%s\r\n", prefix, conf.Address)
			timeOutCtx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.DialTimeout))
			gloConn, _, err = websocket.DefaultDialer.DialContext(timeOutCtx, fmt.Sprintf("%s://%s", prefix, conf.Address), nil)
			if err != nil {
				log.Errorf("连接算力服务器失败:%s\r\n", err.Error())
			} else {
				log.Info("重连算力服务器成功...")
				NewSafeWsConn(conf.HeartbeatIntervalSec)
			}
			//default:

		}
	}
}

func (safeWsConn *SafeWsConn) KeepLive() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Info("发送与算力服务器的保活请求,时间间隔3秒")
			safeWsConn.SendMsg(KeepLiveMsg())
		default:

		}
	}
}

func (safeWsConn *SafeWsConn) Handler() {
	for {
		if gloConn == nil {
			return
		}
		_, msg, err := gloConn.ReadMessage()
		if err != nil {
			log.Error("读取算力服务器消息错误:%s\r\n", err.Error())
			safeWsConn.WatcherCh <- struct{}{}
			//等待重连成功
			return
			//time.Sleep(time.Duration(safeWsConn.HearthIntervalSec + 3))
		}
		msgMap := make(map[string]interface{})
		decodeMsg, _ := safeWsConn.Decode(msg)
		err = json.NewDecoder(bytes.NewReader(decodeMsg)).Decode(&msgMap)
		if err != nil {
			log.Errorf("json decoder msg fail:%s\r\n", err.Error())
			return
		}
		command := msgMap["Command"].(string)
		switch command {
		case "LoginRet":
			//todo 错误处理
			go LoginHandler(decodeMsg)

		case "Resource":
			//todo 错误处理
			go ResourceHandler(decodeMsg)
		case "Warning":
			//todo 错误处理
			go WarningHandler(decodeMsg)

		}
	}
}
